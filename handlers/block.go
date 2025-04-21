package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Block = newJsonHandler(shared.BlockGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonValue("minecraft:block/description/identifier")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.ClientBlock.Get("id"), stores.Feature.Get("block_id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Block.Get("id")
		},
	},
	{
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:block/components/minecraft:geometry"),
			shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
			shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
			shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Geometry.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("geometry_id"), stores.Block.Get("geometry_id"), stores.ClientEntity.Get("geometry_id"))
		},
	},
})
