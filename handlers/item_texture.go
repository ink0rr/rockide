package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var ItemTexture = &JsonHandler{Pattern: shared.ItemTextureGlob}

func init() {
	ItemTexture.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(ClientEntity.Get("spawn_egg"), Item.Get("icon"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ItemTexture.Get("id")
			},
			VanillaData: vanilla.ItemTexture,
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
