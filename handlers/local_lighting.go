package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var LocalLighting = &JsonHandler{
	Pattern: shared.LocalLightingGlob,
	Entries: []JsonEntry{
		{
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonKey("minecraft:local_light_settings/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
	},
}
