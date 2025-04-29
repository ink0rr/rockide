package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Geometry = newJsonHandler(shared.GeometryGlob, []jsonHandlerEntry{
	{
		Path: []shared.JsonPath{
			shared.JsonKey("*"),
			shared.JsonValue("minecraft:geometry/*/description/identifier"),
		},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("geometry_id"), stores.Block.Get("geometry_id"), stores.ClientEntity.Get("geometry_id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Geometry.Get("id")
		},
	},
})
