package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var SpawnRule = &JsonHandler{
	Pattern: shared.SpawnRuleGlob,
	Entries: []JsonEntry{
		{
			Store: stores.EntityId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:spawn_rules/description/identifier"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:permute_type/*/entity_type"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:delay_filter/identifier"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
		},
		{
			Store: stores.BiomeTag.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:biome_filter/**/value")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_biome_tag"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.References.Get()
			},
		},
		{
			Store: stores.EntityEvent.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/event"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/initial_event"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/*/event"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/*/initial_event"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
				if node != nil {
					if id, ok := node.Value.(string); ok {
						return id
					}
				}
				return defaultScope
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
				if node != nil {
					if id, ok := node.Value.(string); ok {
						return stores.EntityEvent.Source.Get(id)
					}
				}
				return stores.EntityEvent.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
				if node != nil {
					if id, ok := node.Value.(string); ok {
						return stores.EntityEvent.References.Get(id)
					}
				}
				return stores.EntityEvent.References.Get()
			},
		},
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_filter"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_prevented_filter"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_filter/*"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_prevented_filter/*"),
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
	},
}
