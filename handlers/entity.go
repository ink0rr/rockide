package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var Entity = newJsonHandler(core.EntityGlob, []jsonHandlerEntry{
	{
		Path:    []string{"minecraft:entity/description/identifier"},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("id")
		},
	},
	{
		Path:       []string{"minecraft:entity/description/animations/*"},
		Actions:    completions | definitions | rename,
		MatchType:  "key",
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate")
		},
	},
	{
		Path:      []string{"minecraft:entity/description/animations/*"},
		Actions:   completions | definitions | rename,
		MatchType: "value",
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("animation_id")
		},
	},
	{
		Path:    []string{"minecraft:entity/description/scripts/animate/*"},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate_refs")
		},
	},
})
