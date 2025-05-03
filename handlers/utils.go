package handlers

import (
	"log"
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
)

func difference(a []core.Symbol, b []core.Symbol) []core.Symbol {
	result := []core.Symbol{}
	set := map[string]bool{}
	for _, ref := range b {
		set[ref.Value] = true
	}
	for _, ref := range a {
		if !set[ref.Value] {
			result = append(result, ref)
		}
	}
	return result
}

func animationControllerSources(id string, stores ...*JsonHandler) []core.Symbol {
	res := []core.Symbol{}
	matchedURIs := make(map[protocol.DocumentURI]bool)
	for _, store := range stores {
		for _, ref := range store.Get("animation_id") {
			if matchedURIs[ref.URI] || ref.Value != id {
				continue
			}
			matchedURIs[ref.URI] = true
			res = append(res, store.GetFrom(ref.URI, "animate")...)
		}
	}
	return res
}

func animationControllerReferences(id string, source *JsonHandler, stores ...*JsonHandler) []core.Symbol {
	res := []core.Symbol{}
	referenceGroup := make(map[protocol.DocumentURI][]core.Symbol)
	for _, store := range stores {
		for _, ref := range store.Get("animation_id") {
			referenceGroup[ref.URI] = append(referenceGroup[ref.URI], ref)
		}
	}
	animationIds := []string{}
	for _, refs := range referenceGroup {
		if !slices.ContainsFunc(refs, func(ref core.Symbol) bool { return ref.Value == id }) {
			continue
		}
		for _, ref := range refs {
			if !slices.Contains(animationIds, ref.Value) {
				animationIds = append(animationIds, ref.Value)
			}
		}
	}
	for _, ref := range source.Get("animate_refs") {
		document, err := textdocument.GetOrReadFile(ref.URI)
		if err != nil {
			log.Println(err)
			continue
		}
		location := jsonc.GetLocation(document.GetText(), document.OffsetAt(ref.Range.Start))
		if id, ok := location.Path[1].(string); ok && slices.Contains(animationIds, id) {
			res = append(res, ref)
		}
	}
	return res
}
