package stores

import "github.com/ink0rr/rockide/shared"

var Recipe = &JsonStore{
	pattern: shared.RecipeGlob,
	entries: []jsonStoreEntry{
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
		},
	},
}
