package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/vanilla"
)

var TerrainTexture = newJsonHandler(shared.TerrainTextureGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Block.Get("texture_id"), stores.ClientBlock.Get("texture_id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.TerrainTexture.Get("id")
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("texture_data/*/textures")},
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Texture.GetPaths()
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
		VanillaData: vanilla.TexturePaths,
	},
})
