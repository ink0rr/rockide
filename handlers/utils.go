package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/stores"
)

func animationControllerSources(id string, stores ...stores.Store) []core.Reference {
	res := []core.Reference{}
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

func animationControllerReferences(id string, source stores.Store, stores ...stores.Store) []core.Reference {
	res := []core.Reference{}
	referenceGroup := make(map[protocol.DocumentURI][]core.Reference)
	for _, store := range stores {
		for _, ref := range store.Get("animation_id") {
			referenceGroup[ref.URI] = append(referenceGroup[ref.URI], ref)
		}
	}
	animationIds := []string{}
	for _, refs := range referenceGroup {
		if !slices.ContainsFunc(refs, func(ref core.Reference) bool { return ref.Value == id }) {
			continue
		}
		for _, ref := range refs {
			if !slices.Contains(animationIds, ref.Value) {
				animationIds = append(animationIds, ref.Value)
			}
		}
	}
	for _, ref := range source.Get("animate_refs") {
		document, _ := textdocument.Open(ref.URI)
		location := jsonc.GetLocation(document.GetText(), document.OffsetAt(&ref.Range.Start))
		if id, ok := location.Path[1].(string); ok && slices.Contains(animationIds, id) {
			res = append(res, ref)
		}
	}
	return res
}
