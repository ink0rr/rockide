package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ItemTexture = &JsonHandler{
	Pattern: shared.ItemTextureGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.ItemTexture.Source,
			Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTexture.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTexture.Source.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("texture_data/*/textures"),
				shared.JsonValue("texture_data/*/textures/*/path"),
			},
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
