package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var LootTable = newJsonHandler(shared.LootTableGlob, []jsonHandlerEntry{
	{
		Path:    []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("id"), stores.Entity.Get("item_id"), stores.Item.Get("item_id"), stores.LootTable.Get("item_id"), stores.Recipe.Get("item_id"))
		},
	},
})
