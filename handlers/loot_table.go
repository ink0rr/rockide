package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var LootTable = &JsonHandler{Pattern: shared.LootTableGlob, SavePath: true}

func init() {
	LootTable.Entries = []JsonEntry{
		{
			Id:   "item_id",
			Path: []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				return entryType != nil && entryType.Value == "item"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Item.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("id"), Entity.Get("item_id"), Item.Get("item_id"), LootTable.Get("item_id"), Recipe.Get("item_id"), TradeTable.Get("item_id"))
			},
			VanillaData: vanilla.ItemIdentifiers,
		},
		{
			Id:   "loot_table_path",
			Path: []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				return entryType != nil && entryType.Value == "loot_table"
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return LootTable.Get("path")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.LootTablePaths,
		},
	}
}
