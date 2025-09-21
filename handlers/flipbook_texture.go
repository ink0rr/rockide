package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var FlipbookTexture = &JsonHandler{
	Pattern: shared.FlipbookTextureGlob,
	Entries: []JsonEntry{
		{
			Store: stores.TerrainTexture.References,
			Path:  []shared.JsonPath{shared.JsonValue("*/atlas_tile")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.References.Get()
			},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("*/flipbook_texture")},
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
