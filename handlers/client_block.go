package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var ClientBlock = &JsonHandler{Pattern: shared.ClientBlockGlob}

func init() {
	ClientBlock.Entries = []JsonEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("*")},
			Matcher: func(ctx *JsonContext) bool {
				return ctx.NodeValue != "format_version"
			},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Block.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ClientBlock.Get("id")
			},
		},
		{
			Id: "texture_id",
			Path: []shared.JsonPath{
				shared.JsonValue("*/textures"),
				shared.JsonValue("*/textures/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return TerrainTexture.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Block.Get("texture_id"), ClientBlock.Get("texture_id"))
			},
			VanillaData: vanilla.TexturePaths,
		},
	}
}
