package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/jsonc"
	"github.com/ink0rr/rockide/stores"
)

var Entity = newJsonHandler(core.EntityGlob, []jsonHandlerEntry{
	{
		Path:       []string{"minecraft:entity/description/identifier"},
		MatchType:  "value",
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
		Path:       []string{"minecraft:entity/description/animations/*"},
		MatchType:  "key",
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
		Path:      []string{"minecraft:entity/description/animations/*"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("animation_id")
		},
	},
	{
		Path:      []string{"minecraft:entity/description/scripts/animate/*/*"},
		MatchType: "key",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:      []string{"minecraft:entity/description/scripts/animate/*"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:       []string{"minecraft:entity/description/properties/*"},
		MatchType:  "key",
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
		Path:      []string{"minecraft:entity/events/**/set_property/*"},
		MatchType: "key",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property_refs")
		},
	},
	{
		Path: []string{
			"minecraft:entity/components/**/filters/**/domain",
			"minecraft:entity/component_groups/**/filters/**/domain",
		},
		MatchType: "value",
		Actions:   completions | definitions | rename,
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
		Path:       []string{"minecraft:entity/component_groups/*"},
		MatchType:  "key",
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
		Path:       []string{"minecraft:entity/events/*"},
		MatchType:  "key",
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
		Path:      []string{"minecraft:entity/events/**/component_groups/*"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group_refs")
		},
	},
	{
		Path: []string{
			"minecraft:entity/components/**/event",
			"minecraft:entity/component_groups/**/event",
		},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "event")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "event_refs")
		},
	},
	{
		Path: []string{
			"minecraft:entity/components/minecraft:type_family/family/*",
			"minecraft:entity/component_groups/*/minecraft:type_family/family/*",
		},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Entity.Get("family"), stores.Entity.Get("family_refs"))
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Path: []string{
			"minecraft:entity/components/**/filters/**/value",
			"minecraft:entity/component_groups/**/filters/**/value",
		},
		MatchType: "value",
		Actions:   completions | definitions | rename,
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
		Path: []string{
			"minecraft:entity/components/minecraft:loot/table",
			"minecraft:entity/component_groups/*/minecraft:loot/table",
		},
		MatchType: "value",
		Actions:   completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.LootTable.Get("path")
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Path: []string{
			"minecraft:entity/components/minecraft:trade_table/table",
			"minecraft:entity/components/minecraft:economy_trade_table/table",
			"minecraft:entity/component_groups/*/minecraft:trade_table/table",
			"minecraft:entity/component_groups/*/minecraft:economy_trade_table/table",
		},
		MatchType: "value",
		Actions:   completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.TradeTable.Get("path")
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
})
