package handlers

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
)

func difference(a []core.Symbol, b []core.Symbol) []core.Symbol {
	res := []core.Symbol{}
	set := mapset.NewThreadUnsafeSet[string]()
	for _, ref := range b {
		set.Add(ref.Value)
	}
	for _, ref := range a {
		if !set.ContainsOne(ref.Value) {
			res = append(res, ref)
		}
	}
	return res
}

func animationControllerSources(id string, stores ...*JsonHandler) []core.Symbol {
	res := []core.Symbol{}
	set := mapset.NewThreadUnsafeSet[protocol.DocumentURI]()
	for _, store := range stores {
		for _, symbol := range store.Get("animation_id", id) {
			if !set.ContainsOne(symbol.URI) {
				set.Add(symbol.URI)
				res = append(res, store.GetFrom(symbol.URI, "animate")...)
			}
		}
	}
	return res
}

func animationControllerReferences(id string, stores ...*JsonHandler) []core.Symbol {
	res := []core.Symbol{}
	set := mapset.NewThreadUnsafeSet[protocol.DocumentURI]()
	for _, store := range stores {
		for _, symbol := range store.Get("animation_id", id) {
			if !set.ContainsOne(symbol.URI) {
				set.Add(symbol.URI)
				res = append(res, store.GetFrom(symbol.URI, "animate_refs")...)
			}
		}
	}
	return res
}
