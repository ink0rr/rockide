package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Geometry = &JsonHandler{
	Pattern: shared.GeometryGlob,
	Entries: []JsonEntry{
		{
			Store: stores.Geometry.Source,
			Path: []shared.JsonPath{
				shared.JsonKey("*"),
				shared.JsonValue("minecraft:geometry/*/description/identifier"),
			},
			Matcher: func(ctx *JsonContext) bool {
				return strings.HasPrefix(ctx.NodeValue, "geometry.")
			},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.Source.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:geometry/*/bones/*/binding"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:geometry/*/description/identifier"),
	},
}
