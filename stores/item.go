package stores

import "github.com/ink0rr/rockide/shared"

var Item = &JsonStore{
	pattern: shared.ItemGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:item/description/identifier")},
		},
		{
			Id: "icon",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:item/components/minecraft:icon"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/texture"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/textures/*"),
			},
		},
		{
			Id:   "tag",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:tags/tags/*")},
		},
		{
			Id:   "item_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/items/*")},
		},
	},
}
