package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var Entity = JsonHandler{pattern: core.EntityGlob, entries: []JsonHandlerEntry{
	{
		Path:    []string{"minecraft:entity/description/identifier"},
		Actions: Completions | Definitions | Rename,
		Source: func(params *JsonParams) []core.Reference {
			return stores.ClientEntity.Get("id")
		},
		References: func(params *JsonParams) []core.Reference {
			return stores.Entity.Get("id")
		},
	},
	{
		Path:    []string{"minecraft:entity/description/animations/*"},
		Actions: Completions | Definitions | Rename,
		Source: func(params *JsonParams) []core.Reference {
			if params.Location.IsAtPropertyKey {
				return stores.Entity.GetFrom(params.URI, "animate")
			}
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
		References: func(params *JsonParams) []core.Reference {
			if params.Location.IsAtPropertyKey {
				return stores.Entity.GetFrom(params.URI, "animation")
			}
			return nil
		},
	},
	{
		Path:    []string{"minecraft:entity/description/scripts/animate/*"},
		Actions: Completions | Definitions | Rename,
		Source: func(params *JsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animation")
		},
		References: func(params *JsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate")
		},
	},
}}
