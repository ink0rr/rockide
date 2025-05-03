package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
)

var biomeTags = [...]string{
	"animal",
	"bamboo",
	"beach",
	"birch",
	"cold",
	"deep",
	"desert",
	"edge",
	"extreme_hills",
	"flower_forest",
	"frozen",
	"hills",
	"ice",
	"ice_plains",
	"jungle",
	"forest",
	"lukewarm",
	"mangrove_swamp",
	"mega",
	"mesa",
	"monster",
	"mooshroom_island",
	"mountain",
	"mutated",
	"nether",
	"no_legacy_worldgen",
	"ocean",
	"overworld",
	"plains",
	"plateau",
	"savanna",
	"swamp",
	"rare",
	"river",
	"roofed",
	"shore",
	"stone",
	"taiga",
	"warm",
	"netherwart_forest",
	"crimson_forest",
	"warped_forest",
	"soulsand_valley",
	"nether_wastes",
	"basalt_deltas",
	"spawn_few_zombified_piglins",
	"spawn_piglin",
	"spawn_endermen",
	"spawn_ghast",
	"spawn_magma_cubes",
	"spawn_many_magma_cubes",
	"sunflower_plains",
}

var equipmentSlots = [...]string{
	"slot.weapon.mainhand",
	"slot.weapon.offhand",
	"slot.armor.head",
	"slot.armor.chest",
	"slot.armor.legs",
	"slot.armor.feet",
	"slot.armor.body",
	"slot.hotbar",
	"slot.inventory",
	"slot.enderchest",
	"slot.saddle",
	"slot.armor",
	"slot.chest",
	"slot.equippable",
}

var inputModes = [...]string{"keyboard_and_mouse", "touch", "gamepad", "motion_controller"}

type molangValue struct {
	references []core.Symbol
	strings    []string
}

var molangTypes = map[string]func() molangValue{
	"BiomeTag": func() molangValue {
		return molangValue{strings: biomeTags[:]}
	},
	"EquipmentSlot": func() molangValue {
		return molangValue{strings: equipmentSlots[:]}
	},
	"InputMode": func() molangValue {
		return molangValue{strings: inputModes[:]}
	},
	"BlockTag": func() molangValue {
		return molangValue{references: Block.Get("tag")}
	},
	"BlockAndItemTag": func() molangValue {
		return molangValue{references: slices.Concat(Block.Get("tag"), Item.Get("tag"))}
	},
	"EntityIdentifier": func() molangValue {
		return molangValue{references: Entity.Get("id")}
	},
	"EntityProperty": func() molangValue {
		return molangValue{references: Entity.Get("property")}
	},
	"TypeFamily": func() molangValue {
		return molangValue{references: Entity.Get("family")}
	},
	"ItemIdentifier": func() molangValue {
		return molangValue{references: Item.Get("id")}
	},
	"ItemTag": func() molangValue {
		return molangValue{references: Item.Get("tag")}
	},
}
