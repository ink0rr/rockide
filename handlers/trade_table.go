package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var TradeTable = &JsonHandler{Pattern: shared.TradeTableGlob, SavePath: true}

func init() {
	TradeTable.Entries = []JsonEntry{
		{
			Id: "item_id",
			Path: []shared.JsonPath{
				shared.JsonValue("tiers/*/groups/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/groups/*/trades/*/wants/*/item"),
				shared.JsonValue("tiers/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/trades/*/wants/*/item"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Item.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("id"), Entity.Get("item_id"), Item.Get("item_id"), LootTable.Get("item_id"), Recipe.Get("item_id"), TradeTable.Get("item_id"))
			},
			VanillaData: vanilla.ItemIdentifiers,
		},
	}
}
