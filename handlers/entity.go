package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Entity = &JsonHandler{
	Pattern: shared.EntityGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.EntityId.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:entity/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
		},
		{
			Store: stores.EntityId.References,
			Path: sliceutil.FlatMap([]string{
				"minecraft:behavior.follow_mob/preferred_actor_type",
				"minecraft:behavior.mingle/mingle_partner_type",
				"minecraft:breedable/breeds_with/baby_type",
				"minecraft:breedable/breeds_with/*/baby_type",
				"minecraft:breedable/breeds_with/mate_type",
				"minecraft:breedable/breeds_with/*/mate_type",
				"minecraft:transformation/into",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
		},
		{
			Store:      stores.Animate.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Animate.References.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Animate.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.Animate.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:entity/description/scripts/animate/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Animate.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Animate.References.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.Animation.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:entity/description/animations/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return ctx.NodeValue
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Animation.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Animation.References.Get()
			},
		},
		{
			Store:      stores.EntityProperty.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/properties/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.References.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.EntityProperty.References,
			Path:  []shared.JsonPath{shared.JsonKey("minecraft:entity/events/**/set_property/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.References.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.EntityProperty.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/domain")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				if test == nil {
					return false
				}
				if value, ok := test.Value.(string); ok {
					return slices.Contains(shared.PropertyTests, value)
				}
				return false
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
				if subject == nil || subject.Value == "self" {
					return stores.EntityProperty.Source.GetFrom(ctx.URI)
				}
				return stores.EntityProperty.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.References.GetFrom(ctx.URI)
			},
		},
		{
			Store:      stores.EntityPropertyValue.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:entity/description/properties/*/values/*")},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return id
				}
				return defaultScope
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return nil
			},
			References: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return nil
			},
		},
		{
			Store: stores.EntityPropertyValue.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:entity/description/properties/*/default")},
			ScopeKey: func(ctx *JsonContext) string {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return id
				}
				return defaultScope
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return nil
			},
			References: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return nil
			},
		},
		{
			Store: stores.EntityPropertyValue.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:entity/events/**/set_property/*")},
			ScopeKey: func(ctx *JsonContext) string {
				path := ctx.GetPath()
				if id, ok := path[len(path)-1].(string); ok {
					return id
				}
				return defaultScope
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				path := ctx.GetPath()
				if id, ok := path[len(path)-1].(string); ok {
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return nil
			},
			References: func(ctx *JsonContext) []core.Symbol {
				path := ctx.GetPath()
				if id, ok := path[len(path)-1].(string); ok {
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return nil
			},
		},
		{
			Store: stores.EntityPropertyValue.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "enum_property"
			},
			ScopeKey: func(ctx *JsonContext) string {
				parent := ctx.GetParentNode()
				domain := jsonc.FindNodeAtLocation(parent, jsonc.Path{"domain"})
				if domain == nil {
					return defaultScope
				}
				if id, ok := domain.Value.(string); ok {
					return id
				}
				return defaultScope
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				domain := jsonc.FindNodeAtLocation(parent, jsonc.Path{"domain"})
				if domain == nil {
					return nil
				}
				id, ok := domain.Value.(string)
				if !ok {
					return nil
				}
				subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
				if subject == nil || subject.Value == "self" {
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return stores.EntityPropertyValue.Source.Get(id)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				domain := jsonc.FindNodeAtLocation(parent, jsonc.Path{"domain"})
				if domain == nil {
					return nil
				}
				id, ok := domain.Value.(string)
				if !ok {
					return nil
				}
				subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
				if subject == nil || subject.Value == "self" {
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return stores.EntityPropertyValue.References.Get(id)
			},
		},
		{
			Store:      stores.EntityComponentGroup.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/component_groups/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.References.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.EntityComponentGroup.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:entity/events/**/component_groups/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.References.GetFrom(ctx.URI)
			},
		},
		{
			Store:      stores.EntityEvent.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/events/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityEvent.References.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityEvent.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.EntityEvent.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/**/event"),
				shared.JsonValue("minecraft:entity/component_groups/**/event"),
				shared.JsonValue("minecraft:entity/events/**/trigger"),
				shared.JsonValue("minecraft:entity/events/**/trigger/event"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityEvent.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityEvent.References.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.EntityEvent.References,
			Path: sliceutil.FlatMap([]string{
				"minecraft:rideable/on_rider_enter_event",
				"minecraft:rideable/on_rider_exit_event",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityEvent.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityEvent.References.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.EntityFamily.Source,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:type_family/family/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:type_family/family/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.EntityFamily.Source.Get(), stores.EntityFamily.References.Get())
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.EntityFamily.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:rideable/family_types/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:rideable/family_types/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityFamily.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityFamily.References.Get()
			},
		},
		{
			Store: stores.EntityFamily.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "is_family"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityFamily.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityFamily.References.Get()
			},
		},
		{
			Store: stores.ItemId.References,
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
			Path: sliceutil.FlatMap([]string{
				"minecraft:ageable/drop_items/*",
				"minecraft:ageable/feed_items",
				"minecraft:ageable/feed_items/*",
				"minecraft:ageable/feed_items/*/item",
				"minecraft:behavior.beg/items/*",
				"minecraft:behavior.charge_held_item/items/*",
				"minecraft:behavior.pickup_items/excluded_items/*",
				"minecraft:behavior.snacking/items/*",
				"minecraft:behavior.tempt/items/*",
				"minecraft:boostable/boost_items/*/item",
				"minecraft:boostable/boost_items/*/replace_item",
				"minecraft:breedable/breed_items",
				"minecraft:breedable/breed_items/*",
				"minecraft:bribeable/bribe_items/*",
				"minecraft:equippable/slots/*/accepted_items/*",
				"minecraft:equippable/slots/*/item",
				"minecraft:giveable/triggers/items/*",
				"minecraft:giveable/triggers/*/items/*",
				"minecraft:healable/items/*/item",
				"minecraft:interact/interactions/transform_to_item",
				"minecraft:interact/interactions/*/transform_to_item",
				"minecraft:item_controllable/control_items",
				"minecraft:item_controllable/control_items/*",
				"minecraft:shareables/items/*/craft_into",
				"minecraft:shareables/items/*/item",
				"minecraft:spawn_entity/entities/spawn_item",
				"minecraft:spawn_entity/entities/*/spawn_item",
				"minecraft:tameable/tame_items",
				"minecraft:tameable/tame_items/*",
				"minecraft:tamemount/auto_reject_items/*/item",
				"minecraft:tamemount/feed_items/*/item",
				"minecraft:trusting/trust_items/*",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
		{
			Store: stores.ItemId.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_equipment"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
		{
			Store: stores.ItemId.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "is_block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Store: stores.ItemTag.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_equipment_tag"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTag.References.Get()
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
			}),
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
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
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.BiomeTag.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
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
			Store: stores.BiomeId.References,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "is_biome"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.References.Get()
			},
		},
	},
	MolangLocations: slices.Concat(
		[]shared.JsonPath{
			shared.JsonValue("minecraft:entity/description/scripts/animate/*/*"),
			shared.JsonValue("minecraft:entity/events/**/set_property/*"),
		},
		sliceutil.FlatMap([]string{
			"minecraft:behavior.eat_block/success_chance",
			"minecraft:experience_reward/on_bred",
			"minecraft:experience_reward/on_death",
			"minecraft:projectile/on_hit/impact_damage/filter",
			"minecraft:rideable/seats/*/rotate_rider_by",
			"minecraft:rideable/seats/rotate_rider_by",
			"minecraft:ambient_sound_interval/event_names/*/condition",
			"minecraft:anger_level/on_increase_sounds/*/condition",
			"minecraft:heartbeat/interval",
		}, func(value string) []shared.JsonPath {
			return []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/" + value),
				shared.JsonValue("minecraft:entity/component_groups/*/" + value),
			}
		}),
	),
}
