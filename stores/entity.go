package stores

import (
	"slices"

	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
)

var Entity = &JsonStore{
	pattern: shared.EntityGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:entity/description/identifier")},
		},
		{
			Id: "id_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:behavior.mingle/mingle_partner_type"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:behavior.mingle/mingle_partner_type"),
			},
		},
		{
			Id:   "animate",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:entity/description/animations/*")},
		},
		{
			Id:   "animation_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:entity/description/animations/*")},
		},
		{
			Id: "animate_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:entity/description/scripts/animate/*/*"),
			},
		},
		{
			Id:   "property",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:entity/description/properties/*")},
		},
		{
			Id:   "property_refs",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:entity/events/**/set_property/*")},
		},
		{
			Id: "property_refs",
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/domain")
			}),
			Transform: func(node *jsonc.Node) *string {
				nodeValue, ok := node.Value.(string)
				if !ok || node.Parent == nil {
					return nil
				}
				parent := node.Parent.Parent
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				if test == nil {
					return nil
				}
				value, ok := test.Value.(string)
				if ok && slices.Contains(shared.PropertyTests, value) {
					return &nodeValue
				}
				return nil
			},
		},
		{
			Id:   "component_group",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:entity/component_groups/*")},
		},
		{
			Id:   "component_group_refs",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:entity/events/**/component_groups/*")},
		},
		{
			Id:   "event",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:entity/events/*")},
		},
		{
			Id: "event_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/**/event"),
				shared.JsonValue("minecraft:entity/component_groups/**/event"),
				shared.JsonValue("minecraft:entity/events/**/trigger"),
				shared.JsonValue("minecraft:entity/events/**/trigger/event"),
			},
		},
		{
			Id: "family",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:type_family/family/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:type_family/family/*"),
			},
		},
		{
			Id: "family_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:rideable/family_types/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:rideable/family_types/*"),
			},
		},
		{
			Id: "family_refs",
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Transform: func(node *jsonc.Node) *string {
				nodeValue, ok := node.Value.(string)
				if !ok || node.Parent == nil {
					return nil
				}
				parent := node.Parent.Parent
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				if test != nil && test.Value == "is_family" {
					return &nodeValue
				}
				return nil
			},
		},
		{
			Id: "block_id",
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
		},
		{
			Id: "item_id",
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
		},
		{
			Id: "item_id",
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Transform: func(node *jsonc.Node) *string {
				nodeValue, ok := node.Value.(string)
				if !ok || node.Parent == nil {
					return nil
				}
				parent := node.Parent.Parent
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				if test == nil {
					return nil
				}
				testValue, ok := test.Value.(string)
				if !ok || testValue != "has_equipment" {
					return nil
				}
				return &nodeValue
			},
		},
		{
			Id: "loot_table_path",
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
		},
		{
			Id: "trade_table_path",
			Path: sliceutil.FlatMap([]string{
				"minecraft:trade_table/table",
				"minecraft:economy_trade_table/table",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
		},
	},
}
