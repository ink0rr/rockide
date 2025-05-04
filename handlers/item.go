package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var Item = &JsonHandler{Pattern: shared.ItemGlob}

func init() {
	Item.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:item/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("id"), Entity.Get("item_id"), Item.Get("item_id"), LootTable.Get("item_id"), Recipe.Get("item_id"), TradeTable.Get("item_id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Item.Get("id")
			},
		},
		{
			Id: "icon",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:item/components/minecraft:icon"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/texture"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/textures/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return ItemTexture.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Item.Get("icon"), ClientEntity.Get("spawn_egg"))
			},
			VanillaData: vanilla.ItemTexture,
		},
		{
			Id:   "tag",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:tags/tags/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Item.Get("tag")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.ItemTag,
		},
		{
			Id:   "item_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/items/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Block.Get("id"), Item.Get("id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("id"), ClientBlock.Get("id"), Entity.Get("item_id"), Item.Get("item_id"), LootTable.Get("item_id"), Recipe.Get("item_id"), TradeTable.Get("item_id"))
			},
			VanillaData: vanilla.ItemIdentifiers,
		},
	}
}
