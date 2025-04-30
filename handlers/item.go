package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/vanilla"
)

var Item = newJsonHandler(shared.ItemGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonValue("minecraft:item/description/identifier")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("id"), stores.Entity.Get("item_id"), stores.Item.Get("item_id"), stores.LootTable.Get("item_id"), stores.Recipe.Get("item_id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
	},
	{
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:item/components/minecraft:icon"),
			shared.JsonValue("minecraft:item/components/minecraft:icon/texture"),
			shared.JsonValue("minecraft:item/components/minecraft:icon/textures/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ItemTexture.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Item.Get("icon"), stores.ClientEntity.Get("spawn_egg"))
		},
		VanillaData: vanilla.ItemTexture,
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/items/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("id"), stores.Entity.Get("item_id"), stores.Item.Get("item_id"), stores.LootTable.Get("item_id"), stores.Recipe.Get("item_id"))
		},
		VanillaData: vanilla.ItemIdentifiers,
	},
})
