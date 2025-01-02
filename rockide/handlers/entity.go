package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/rockide/core"
	"github.com/ink0rr/rockide/rockide/stores"
)

var Entity = NewJsonHandler(core.EntityGlob, []*JsonHandlerEntry{
	{
		path: []string{"minecraft:entity/description/identifier"},
		Completions: func(params *JsonHandlerParams) []core.Reference {
			return stores.Difference(stores.ClientEntity.Get("id"), stores.Entity.Get("id"))
		},
	},
	{
		path: []string{"minecraft:entity/description/animations/*"},
		Completions: func(params *JsonHandlerParams) []core.Reference {
			if params.IsAtPropertyKeyOrArray() {
				return stores.Difference(
					stores.Entity.GetFrom(params.URI, "animate"),
					stores.Entity.GetFrom(params.URI, "animation"),
				)
			}
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
		Definitions: func(params *JsonHandlerParams) []core.Reference {
			if params.IsAtPropertyKeyOrArray() {
				return stores.Entity.GetFrom(params.URI, "animate")
			}
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
		Rename: func(params *JsonHandlerParams) []core.Reference {
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
	},
})
