package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var Entity = &JsonHandler{Pattern: shared.EntityGlob}

func init() {
	Entity.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:entity/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(ClientEntity.Get("id"), Entity.Get("id_refs"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.Get("id")
			},
		},
		{
			Id: "id_refs",
			Path: sliceutil.FlatMap([]string{
				"minecraft:behavior.mingle/mingle_partner_type",
				"minecraft:breedable/breeds_with/baby_type",
				"minecraft:breedable/breeds_with/mate_type",
				"minecraft:transformation/into",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(ClientEntity.Get("id"), Entity.Get("id_refs"))
			},
			VanillaData: vanilla.EntityIdentifiers,
		},
		{
			Id:         "animate",
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "animate_refs")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "animate")
			},
		},
		{
			Id:   "animation_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:entity/description/animations/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(AnimationController.Get("id"), Animation.Get("id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.Get("animation_id")
			},
		},
		{
			Id: "animate_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:entity/description/scripts/animate/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "animate")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "animate_refs")
			},
		},
		{
			Id:         "property",
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/properties/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "property_refs")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "property")
			},
		},
		{
			Id:   "property_refs",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:entity/events/**/set_property/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "property")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "property_refs")
			},
		},
		{
			Id: "property_refs",
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
					return Entity.GetFrom(ctx.URI, "property")
				}
				return Entity.Get("property")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "property_refs")
			},
		},
		{
			Id:         "component_group",
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/component_groups/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "component_group_refs")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "component_group")
			},
		},
		{
			Id:   "component_group_refs",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:entity/events/**/component_groups/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "component_group")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "component_group_refs")
			},
		},
		{
			Id:         "event",
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/events/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "event_refs")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "event")
			},
		},
		{
			Id: "event_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/**/event"),
				shared.JsonValue("minecraft:entity/component_groups/**/event"),
				shared.JsonValue("minecraft:entity/events/**/trigger"),
				shared.JsonValue("minecraft:entity/events/**/trigger/event"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "event")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.GetFrom(ctx.URI, "event_refs")
			},
		},
		{
			Id: "family",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:type_family/family/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:type_family/family/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Entity.Get("family"), Entity.Get("family_refs"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.Family,
		},
		{
			Id: "family_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:rideable/family_types/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:rideable/family_types/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.Get("family")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.Get("family_refs")
			},
			VanillaData: vanilla.Family,
		},
		{
			Id: "family_refs",
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "is_family"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Entity.Get("family")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Entity.Get("family_refs")
			},
			VanillaData: vanilla.Family,
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
			Source: func(ctx *JsonContext) []core.Symbol {
				return Block.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Entity.Get("block_id"), Feature.Get("block_id"))
			},
			VanillaData: vanilla.BlockIdentifiers,
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
			Source: func(ctx *JsonContext) []core.Symbol {
				return Item.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("id"), Entity.Get("item_id"), Item.Get("item_id"), LootTable.Get("item_id"), Recipe.Get("item_id"), TradeTable.Get("item_id"))
			},
			VanillaData: vanilla.ItemIdentifiers,
		},
		{
			Id: "item_id",
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_equipment"
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
			}),
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return LootTable.Get("path")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.LootTablePaths,
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
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return TradeTable.Get("path")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.TradeTablePaths,
		},
	}
	Entity.MolangLocations = slices.Concat(
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
	)
}
