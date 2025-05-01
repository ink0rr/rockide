package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/vanilla"
)

var LootTable = newJsonHandler(shared.LootTableGlob, []jsonHandlerEntry{
	{
		Path: []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
		Matcher: func(params *jsonParams) bool {
			parent := params.getParentNode()
			entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
			return entryType != nil && entryType.Value == "item"
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("id"), stores.Entity.Get("item_id"), stores.Item.Get("item_id"), stores.LootTable.Get("item_id"), stores.Recipe.Get("item_id"))
		},
		VanillaData: vanilla.ItemIdentifiers,
	},
	{
		Path: []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
		Matcher: func(params *jsonParams) bool {
			parent := params.getParentNode()
			entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
			return entryType != nil && entryType.Value == "loot_table"
		},
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.LootTable.Get("path")
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
		VanillaData: vanilla.LootTablePaths,
	},
})
