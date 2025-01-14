package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/stores"
)

var Entity = newJsonHandler(core.EntityGlob, []jsonHandlerEntry{
	{
		Matcher:    []jsonPath{matchValue("minecraft:entity/description/identifier")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("id")
		},
	},
	{
		Matcher:    []jsonPath{matchKey("minecraft:entity/description/animations/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate")
		},
	},
	{
		Matcher: []jsonPath{matchValue("minecraft:entity/description/animations/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("animation_id")
		},
	},
	{
		Matcher: []jsonPath{
			matchKey("minecraft:entity/description/scripts/animate/*/*"),
			matchValue("minecraft:entity/description/scripts/animate/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Matcher:    []jsonPath{matchKey("minecraft:entity/description/properties/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property")
		},
	},
	{
		Matcher: []jsonPath{matchKey("minecraft:entity/events/**/set_property/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property_refs")
		},
	},
	{
		Matcher: []jsonPath{
			matchValue("minecraft:entity/components/**/filters/**/domain"),
			matchValue("minecraft:entity/component_groups/**/filters/**/domain"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			parent := params.getParentNode()
			test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
			if test == nil {
				return nil
			}
			if value, ok := test.Value.(string); !ok || !slices.Contains(core.PropertyDomain, value) {
				return nil
			}
			subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
			if subject == nil || subject.Value == "self" {
				return stores.Entity.GetFrom(params.URI, "property")
			}
			return stores.Entity.Get("property")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property_refs")
		},
	},
	{
		Matcher:    []jsonPath{matchKey("minecraft:entity/component_groups/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group")
		},
	},
	{
		Matcher:    []jsonPath{matchKey("minecraft:entity/events/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "event_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "event")
		},
	},
	{
		Matcher: []jsonPath{matchValue("minecraft:entity/events/**/component_groups/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group_refs")
		},
	},
	{
		Matcher: []jsonPath{
			matchValue("minecraft:entity/components/**/event"),
			matchValue("minecraft:entity/component_groups/**/event"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "event")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "event_refs")
		},
	},
	{
		Matcher: []jsonPath{
			matchValue("minecraft:entity/components/minecraft:type_family/family/*"),
			matchValue("minecraft:entity/component_groups/*/minecraft:type_family/family/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Entity.Get("family"), stores.Entity.Get("family_refs"))
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Matcher: []jsonPath{
			matchValue("minecraft:entity/components/**/filters/**/value"),
			matchValue("minecraft:entity/component_groups/**/filters/**/value"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			parent := params.getParentNode()
			test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
			if test == nil || test.Value != "is_family" {
				return nil
			}
			return stores.Entity.Get("family")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("family_refs")
		},
	},
	{
		Matcher: []jsonPath{
			matchValue("minecraft:entity/components/minecraft:loot/table"),
			matchValue("minecraft:entity/component_groups/*/minecraft:loot/table"),
		},
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.LootTable.Get("path")
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Matcher: []jsonPath{
			matchValue("minecraft:entity/components/minecraft:trade_table/table"),
			matchValue("minecraft:entity/components/minecraft:economy_trade_table/table"),
			matchValue("minecraft:entity/component_groups/*/minecraft:trade_table/table"),
			matchValue("minecraft:entity/component_groups/*/minecraft:economy_trade_table/table"),
		},
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.TradeTable.Get("path")
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
})
