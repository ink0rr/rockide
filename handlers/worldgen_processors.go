package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var WorldgenProcessor = &JsonHandler{
	Pattern: shared.WorldgenProcessorGlob,
	Entries: []JsonEntry{
		{
			Store: stores.WorldgenProcessor.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:processor_list/description/identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenProcessor.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenProcessor.Source.Get()
			},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:processor_list/processors/**/rules/*/block_entity_modifier/loot_table")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block_state"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block_state/name"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block_state"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block_state/name"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/output_state"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/output_state/name"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Store: stores.BlockTag.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/tag"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/tag"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockTag.References.Get()
			},
		},
		// TODO: Block states
	},
}
