package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var LootTable = &JsonHandler{
	Pattern:   shared.LootTableGlob,
	PathStore: stores.LootTablePath,
	Entries: []JsonEntry{
		{
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				return entryType != nil && entryType.Value == "item"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				return entryType != nil && entryType.Value == "loot_table"
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
