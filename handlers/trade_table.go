package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var TradeTable = &JsonHandler{
	Pattern:   shared.TradeTableGlob,
	PathStore: stores.TradeTablePath,
	Entries: []JsonEntry{
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("tiers/*/groups/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/groups/*/trades/*/wants/*/item"),
				shared.JsonValue("tiers/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/trades/*/wants/*/item"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
	},
}
