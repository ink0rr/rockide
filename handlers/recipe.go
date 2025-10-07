package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Recipe = &JsonHandler{
	Pattern: shared.RecipeGlob,
	Entries: []JsonEntry{
		{
			Store: stores.ItemId.References,
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
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
		{
			Store: stores.RecipeTag.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:recipe_shaped/tags/*"),
				shared.JsonValue("minecraft:recipe_shapeless/tags/*"),
				shared.JsonValue("minecraft:recipe_furnace/tags"),
				shared.JsonValue("minecraft:recipe_brewing_container/tags"),
				shared.JsonValue("minecraft:recipe_brewing_container_mix/tags"),
				shared.JsonValue("minecraft:recipe_smithing_transform/tags"),
				shared.JsonValue("minecraft:recipe_smithing_trim/tags"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.RecipeTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.RecipeTag.References.Get()
			},
		},
		{
			Store: stores.ItemTag.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:recipe_shaped/key/*/tag"),
				shared.JsonValue("minecraft:recipe_shapeless/ingredients/*/tag"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTag.References.Get()
			},
		},
	},
}
