package stores

import "github.com/ink0rr/rockide/shared"

var Item = newJsonStore(shared.ItemGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:item/description/identifier"},
	},
	{
		Id: "icon",
		Path: []string{
			"minecraft:item/components/minecraft:icon",
			"minecraft:item/components/minecraft:icon/texture",
			"minecraft:item/components/minecraft:icon/textures/*",
		},
		Transform: skipKey,
	},
	{
		Id:   "tag",
		Path: []string{"minecraft:item/components/minecraft:tags/tags"},
	},
	{
		Id:   "item_id",
		Path: []string{"minecraft:item/components/minecraft:repairable/repair_items/*/items"},
	},
})
