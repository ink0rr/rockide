package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientBlock = &JsonHandler{
	Pattern: shared.ClientBlockGlob,
	Entries: []JsonEntry{
		{
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonKey("*")},
			Matcher: func(ctx *JsonContext) bool {
				return ctx.NodeValue != "format_version"
			},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Store: stores.TerrainTexture.References,
			Path: []shared.JsonPath{
				shared.JsonValue("*/textures"),
				shared.JsonValue("*/textures/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.References.Get()
			},
		},
	},
}
