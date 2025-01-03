package stores

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/jsonc"
)

var Entity = newJsonStore(core.EntityGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:entity/description/identifier"},
	},
	{
		Id:   "animation",
		Path: []string{"minecraft:entity/description/animations"},
	},
	{
		Id:   "animation_id",
		Path: []string{"minecraft:entity/description/animations/*"},
	},
	{
		Id:   "animate",
		Path: []string{"minecraft:entity/description/scripts/animate"},
	},
	{
		Id:   "property",
		Path: []string{"minecraft:entity/description/properties"},
	},
	{
		Id:   "property_refs",
		Path: []string{"minecraft:entity/events/**/set_property"},
	},
	{
		Id: "property_refs",
		Path: []string{
			"minecraft:entity/components/**/filters/**/domain",
			"minecraft:entity/component_groups/**/filters/**/domain",
		},
		Transform: func(node *jsonc.Node) transformResult {
			nodeValue, ok := node.Value.(string)
			if !ok || node.Parent == nil {
				return transformResult{Skip: true}
			}
			parent := node.Parent.Parent
			test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
			if test == nil {
				return transformResult{Skip: true}
			}
			testValue, ok := test.Value.(string)
			if !ok || slices.Index(core.PropertyDomain, testValue) == -1 {
				return transformResult{Skip: true}
			}
			return transformResult{Value: nodeValue}
		},
	},
	{
		Id:   "component_group",
		Path: []string{"minecraft:entity/component_groups"},
	},
	{
		Id:   "component_group_refs",
		Path: []string{"minecraft:entity/events/**/component_groups"},
	},
	{
		Id:   "event",
		Path: []string{"minecraft:entity/events"},
	},
	{
		Id:        "event_refs",
		Path:      []string{"minecraft:entity/events/**/trigger", "minecraft:entity/events/**/trigger/event"},
		Transform: skipKey,
	},
	{
		Id: "event_refs",
		Path: []string{
			"minecraft:behavior.admire_item",
			"minecraft:behavior.avoid_block",
			"minecraft:behavior.avoid_mob_type",
			"minecraft:behavior.celebrate_survive",
			"minecraft:behavior.celebrate",
			"minecraft:behavior.defend_trusted_target",
			"minecraft:behavior.delayed_attack",
			"minecraft:behavior.dig",
			"minecraft:behavior.drop_item_for",
			"minecraft:behavior.eat_block",
			"minecraft:behavior.emerge",
			"minecraft:behavior.go_and_give_items_to_noteblock",
			"minecraft:behavior.go_and_give_items_to_owner",
			"minecraft:behavior.go_home",
			"minecraft:behavior.hold_ground",
			"minecraft:behavior.knockback_roar",
			"minecraft:behavior.lay_egg",
			"minecraft:behavior.melee_attack",
			"minecraft:behavior.stomp_attack",
			"minecraft:behavior.move_to_block",
			"minecraft:behavior.ram_attack",
			"minecraft:behavior.work",
			"minecraft:behavior.work_composter",
			"minecraft:ageable",
			"minecraft:angry",
			"minecraft:attack_cooldown",
			"minecraft:block_sensor",
			"minecraft:breedable",
			"minecraft:damage_sensor",
			"minecraft:environment_sensor",
			"minecraft:drying_out_timer",
			"minecraft:equippable",
			"minecraft:genetics",
			"minecraft:giveable",
			"minecraft:inside_block_notifier",
			"minecraft:interact",
			"minecraft:leashable",
			"minecraft:lookat",
			"minecraft:nameable",
			"minecraft:on_death",
			"minecraft:on_friendly_anger",
			"minecraft:on_hurt",
			"minecraft:on_hurt_by_player",
			"minecraft:on_ignite",
			"minecraft:on_start_landing",
			"minecraft:on_start_takeoff",
			"minecraft:on_target_acquired",
			"minecraft:on_target_escape",
			"minecraft:on_wake_with_owner",
			"minecraft:peek",
			"minecraft:projectile",
			"minecraft:raid_trigger",
			"minecraft:rail_sensor",
			"minecraft:ravager_blocked",
			"minecraft:scheduler",
			"minecraft:sittable",
			"minecraft:tameable",
			"minecraft:tamemount",
			"minecraft:target_nearby_sensor",
			"minecraft:timer",
			"minecraft:trusting",
		},
	},
	{
		Id: "family",
		Path: []string{
			"minecraft:entity/components/minecraft:type_family/family",
			"minecraft:entity/component_groups/*/minecraft:type_family/family",
		},
	},
	{
		Id:   "family_refs",
		Path: []string{"minecraft:entity/components/**/filters/**/value", "minecraft:entity/component_groups/**/filters/**/value"},
		Transform: func(node *jsonc.Node) transformResult {
			nodeValue, ok := node.Value.(string)
			if !ok || node.Parent == nil {
				return transformResult{Skip: true}
			}
			parent := node.Parent.Parent
			test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
			if test == nil {
				return transformResult{Skip: true}
			}
			testValue, ok := test.Value.(string)
			if !ok || testValue != "is_family" {
				return transformResult{Skip: true}
			}
			return transformResult{Value: nodeValue}
		},
	},
	{
		Id: "loot_table_path",
		Path: []string{
			"minecraft:loot/table",
			"minecraft:behavior.sneeze/loot_table",
			"minecraft:barter/barter_table",
			"minecraft:interact/interactions/add_items/table",
			"minecraft:interact/interactions/*/add_items/table",
			"minecraft:interact/interactions/spawn_items/table",
			"minecraft:interact/interactions/*/spawn_items/table",
		},
	},
	{
		Id:   "trade_table_path",
		Path: []string{"minecraft:trade_table/table", "minecraft:economy_trade_table/table"},
	},
})
