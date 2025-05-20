package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
)

var Geometry = &JsonHandler{Pattern: shared.GeometryGlob}

func init() {
	Geometry.Entries = []JsonEntry{
		{
			Id: "id",
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
				return slices.Concat(Attachable.Get("geometry_id"), Block.Get("geometry_id"), ClientEntity.Get("geometry_id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Geometry.Get("id")
			},
		},
	}
	Geometry.MolangLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:geometry/*/bones/*/binding"),
	}
	Geometry.MolangSemanticLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:geometry/*/description/identifier"),
	}
}
