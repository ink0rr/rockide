package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var TerrainTexture = newJsonHandler(shared.TerrainTextureGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientBlock.Get("texture")
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
	},
})
