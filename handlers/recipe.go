package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var Recipe = &JsonHandler{Pattern: shared.RecipeGlob}

func init() {
	Recipe.Entries = []JsonEntry{
		{
			Id: "item_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:recipe_furnace/input"),
				shared.JsonValue("minecraft:recipe_furnace/output"),
				shared.JsonValue("minecraft:recipe_shaped/result/item"),
				shared.JsonValue("minecraft:recipe_shaped/key/*/item"),
				shared.JsonValue("minecraft:recipe_shapeless/result/item"),
				shared.JsonValue("minecraft:recipe_shapeless/ingredients/*/item"),
				shared.JsonValue("minecraft:recipe_brewing_mix/input"),
				shared.JsonValue("minecraft:recipe_brewing_mix/reagent"),
				shared.JsonValue("minecraft:recipe_brewing_mix/output"),
				shared.JsonValue("minecraft:recipe_brewing_container/input"),
				shared.JsonValue("minecraft:recipe_brewing_container/reagent"),
				shared.JsonValue("minecraft:recipe_brewing_container/output"),
			},
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
