package stores

import "github.com/ink0rr/rockide/core"

var Recipe = newJsonStore(core.RecipeGlob, []jsonStoreEntry{
	{
		Id: "item_id",
		Path: []string{
			"minecraft:recipe_furnace/input",
			"minecraft:recipe_furnace/output",

			"minecraft:recipe_shaped/result/item",
			"minecraft:recipe_shaped/key/*/item",

			"minecraft:recipe_shapeless/result/item",
			"minecraft:recipe_shapeless/ingredients/*/item",

			"minecraft:recipe_brewing_mix/input",
			"minecraft:recipe_brewing_mix/reagent",
			"minecraft:recipe_brewing_mix/output",

			"minecraft:recipe_brewing_container/input",
			"minecraft:recipe_brewing_container/reagent",
			"minecraft:recipe_brewing_container/output",
		},
	},
})
