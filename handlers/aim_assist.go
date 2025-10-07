package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var AimAssistPreset = &JsonHandler{
	Pattern: shared.AimAssistPresetGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.AimAssistId.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:aim_assist_preset/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistId.Source.Get()
			},
		},
		{
			Store: stores.AimAssistCategory.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:aim_assist_preset/default_item_settings"),
				shared.JsonValue("minecraft:aim_assist_preset/hand_settings"),
				shared.JsonValue("minecraft:aim_assist_preset/item_settings/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistCategory.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistCategory.References.Get()
			},
		},
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:aim_assist_preset/exclusion_list/*"),
				shared.JsonValue("minecraft:aim_assist_preset/liquid_targeting_list/*"),
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
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonKey("minecraft:aim_assist_preset/item_settings/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
	},
}

var AimAssistCategory = &JsonHandler{
	Pattern: shared.AimAssistCategoryGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.AimAssistCategory.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:aim_assist_categories/categories/*/name")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistCategory.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistCategory.Source.Get()
			},
		},
		{
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonKey("minecraft:aim_assist_categories/categories/*/priorities/blocks/*")},
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
			Store: stores.EntityId.References,
			Path:  []shared.JsonPath{shared.JsonKey("minecraft:aim_assist_categories/categories/*/priorities/entities/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
		},
	},
}
