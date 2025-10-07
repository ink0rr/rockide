package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Atmosphere = &JsonHandler{
	Pattern: shared.AtmosphereGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.Atmosphere.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:atmosphere_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Atmosphere.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Atmosphere.Source.Get()
			},
		},
	},
}
