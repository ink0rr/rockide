package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Fog = &JsonHandler{
	Pattern: shared.FogGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.Fog.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:fog_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Fog.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Fog.Source.Get()
			},
		},
	},
}
