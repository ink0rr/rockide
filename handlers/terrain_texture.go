package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var TerrainTexture = &JsonHandler{Pattern: shared.TerrainTextureGlob}

func init() {
	TerrainTexture.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Block.Get("texture_id"), ClientBlock.Get("texture_id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return TerrainTexture.Get("id")
			},
		},
		{
			Id:            "texture_path",
			Path:          []shared.JsonPath{shared.JsonValue("texture_data/*/textures")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Texture.GetPaths()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.TexturePaths,
		},
	}
}
