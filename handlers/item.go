package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var Item = newJsonHandler(core.ItemGlob, []jsonHandlerEntry{
	{
		Matcher:    []jsonPath{matchValue("minecraft:item/description/identifier")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Item.Get("item_id"), stores.Attachable.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
	},
	{
		Matcher: []jsonPath{
			matchValue("minecraft:item/components/minecraft:icon"),
			matchValue("minecraft:item/components/minecraft:icon/texture"),
			matchValue("minecraft:item/components/minecraft:icon/textures/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ItemTexture.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Item.Get("icon"), stores.ClientEntity.Get("spawn_egg"))
		},
	},
	{
		Matcher: []jsonPath{matchValue("minecraft:item/components/minecraft:repairable/repair_items/*/items/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("item_id")
		},
	},
})
