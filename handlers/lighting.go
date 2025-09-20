package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Lighting = &JsonHandler{
	Pattern: shared.LightingGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.Lighting.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:lighting_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Lighting.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Lighting.Source.Get()
			},
		},
	},
}
