package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Entity = newJsonHandler(shared.EntityGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonValue("minecraft:entity/description/identifier")},
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
		Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/animations/*")},
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
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:entity/description/animations/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.AnimationController.Get("id"), stores.Animation.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("animation_id")
		},
	},
	{
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:entity/description/scripts/animate/*"),
			shared.JsonKey("minecraft:entity/description/scripts/animate/*/*"),
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
		Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/properties/*")},
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
		Path:    []shared.JsonPath{shared.JsonKey("minecraft:entity/events/**/set_property/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "property_refs")
		},
	},
	{
		Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
			return shared.JsonValue(path + "/domain")
		}),
		Matcher: func(params *jsonParams) bool {
			parent := params.getParentNode()
			test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
			if test == nil {
				return false
			}
			if value, ok := test.Value.(string); ok {
				return slices.Contains(shared.PropertyTests, value)
			}
			return false
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			parent := params.getParentNode()
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
		Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/component_groups/*")},
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
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:entity/events/**/component_groups/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.GetFrom(params.URI, "component_group_refs")
		},
	},
	{
		Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/events/*")},
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
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:entity/components/**/event"),
			shared.JsonValue("minecraft:entity/component_groups/**/event"),
			shared.JsonValue("minecraft:entity/events/**/trigger"),
			shared.JsonValue("minecraft:entity/events/**/trigger/event"),
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
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:entity/components/minecraft:type_family/family/*"),
			shared.JsonValue("minecraft:entity/component_groups/*/minecraft:type_family/family/*"),
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
		Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
			return shared.JsonValue(path + "/value")
		}),
		Matcher: func(params *jsonParams) bool {
			parent := params.getParentNode()
			test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
			return test != nil && test.Value == "is_family"
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("family")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("family_refs")
		},
	},
	{
		Path: sliceutil.FlatMap([]string{
			"minecraft:behavior.avoid_block/target_blocks",
			"minecraft:behavior.eat_block/eat_and_replace_block_pairs/*/eat_block",
			"minecraft:behavior.eat_block/eat_and_replace_block_pairs/*/replace_block",
			"minecraft:behavior.jump_to_block/forbidden_blocks",
			"minecraft:behavior.jump_to_block/preferred_blocks",
			"minecraft:behavior.lay_egg/egg_type",
			"minecraft:behavior.lay_egg/target_blocks",
			"minecraft:behavior.move_to_block/target_blocks",
			"minecraft:behavior.raid_garden/blocks",
			"minecraft:behavior.random_search_and_dig/target_blocks",
			"minecraft:block_sensor/on_break/*/block_list",
			"minecraft:break_blocks/breakable_blocks",
			"minecraft:breathable/breathe_blocks",
			"minecraft:breathable/non_breathe_blocks",
			"minecraft:breedable/environment_requirements/blocks",
			"minecraft:breedable/environment_requirements/*/blocks",
			"minecraft:buoyant/liquid_blocks",
			"minecraft:home/home_block_list",
			"minecraft:inside_block_notifier/block_list/*/block/name",
			"minecraft:navigation.climb/blocks_to_avoid",
			"minecraft:navigation.float/blocks_to_avoid",
			"minecraft:navigation.fly/blocks_to_avoid",
			"minecraft:navigation.generic/blocks_to_avoid",
			"minecraft:navigation.hover/blocks_to_avoid",
			"minecraft:navigation.swim/blocks_to_avoid",
			"minecraft:navigation.walk/blocks_to_avoid",
			"minecraft:preferred_path/preferred_path_blocks/blocks",
			"minecraft:preferred_path/preferred_path_blocks/blocks/*/name",
			"minecraft:trail/block_type",
			"minecraft:transformation/delay/block_types",
		}, func(value string) []shared.JsonPath {
			return []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/" + value),
				shared.JsonValue("minecraft:entity/component_groups/*/" + value),
			}
		}),
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Block.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Entity.Get("block_id"), stores.Feature.Get("block_id"))
		},
	},
	{
		Path: sliceutil.FlatMap([]string{
			"minecraft:ageable/drop_items",
			"minecraft:ageable/feed_items",
			"minecraft:ageable/feed_items/*/item",
			"minecraft:behavior.beg/items",
			"minecraft:behavior.charge_held_item/items",
			"minecraft:behavior.pickup_items/excluded_items",
			"minecraft:behavior.snacking/items",
			"minecraft:behavior.tempt/items",
			"minecraft:boostable/boost_items/*/item",
			"minecraft:boostable/boost_items/*/replace_item",
			"minecraft:breedable/breed_items",
			"minecraft:bribeable/bribe_items",
			"minecraft:equippable/slots/*/accepted_items",
			"minecraft:equippable/slots/*/item",
			"minecraft:giveable/triggers/items",
			"minecraft:giveable/triggers/*/items",
			"minecraft:healable/items/*/item",
			"minecraft:interact/interactions/transform_to_item",
			"minecraft:interact/interactions/*/transform_to_item",
			"minecraft:item_controllable/control_items",
			"minecraft:shareables/items/*/craft_into",
			"minecraft:shareables/items/*/item",
			"minecraft:spawn_entity/entities/spawn_item",
			"minecraft:spawn_entity/entities/*/spawn_item",
			"minecraft:tameable/tame_items",
			"minecraft:tamemount/auto_reject_items/*/item",
			"minecraft:tamemount/feed_items/*/item",
			"minecraft:trusting/trust_items",
		}, func(value string) []shared.JsonPath {
			return []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/" + value),
				shared.JsonValue("minecraft:entity/component_groups/*/" + value),
			}
		}),
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("id"), stores.Entity.Get("item_id"), stores.Item.Get("item_id"), stores.LootTable.Get("item_id"), stores.Recipe.Get("item_id"))
		},
	},
	{
		Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
			return shared.JsonValue(path + "/value")
		}),
		Matcher: func(params *jsonParams) bool {
			parent := params.getParentNode()
			test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
			return test != nil && test.Value == "has_equipment"
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("id"), stores.Entity.Get("item_id"), stores.Item.Get("item_id"), stores.LootTable.Get("item_id"), stores.Recipe.Get("item_id"))
		},
	},
	{
		Path: sliceutil.FlatMap([]string{
			"minecraft:loot/table",
			"minecraft:behavior.sneeze/loot_table",
			"minecraft:barter/barter_table",
			"minecraft:interact/interactions/add_items/table",
			"minecraft:interact/interactions/*/add_items/table",
			"minecraft:interact/interactions/spawn_items/table",
			"minecraft:interact/interactions/*/spawn_items/table",
		}, func(value string) []shared.JsonPath {
			return []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/" + value),
				shared.JsonValue("minecraft:entity/component_groups/*/" + value),
			}
		},
		),
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.LootTable.Get("path")
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Path: sliceutil.FlatMap([]string{
			"minecraft:trade_table/table",
			"minecraft:economy_trade_table/table",
		}, func(value string) []shared.JsonPath {
			return []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/" + value),
				shared.JsonValue("minecraft:entity/component_groups/*/" + value),
			}
		}),
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.TradeTable.GetPaths()
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
})
