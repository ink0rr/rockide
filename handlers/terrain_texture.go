package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var TerrainTexture = &JsonHandler{
	Pattern: shared.TerrainTextureGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.TerrainTexture.Source,
			Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.Source.Get()
			},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("texture_data/*/textures")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
