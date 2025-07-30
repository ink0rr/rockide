package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/vanilla"
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

var graphicsModes = [...]string{"simple", "fancy", "deferred", "raytraced"}

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
	"GraphicsMode": func() molangValue {
		return molangValue{strings: graphicsModes[:]}
	},
	"InputMode": func() molangValue {
		return molangValue{strings: inputModes[:]}
	},
	"BlockTag": func() molangValue {
		return molangValue{references: stores.BlockTag.Source.Get()}
	},
	"BlockAndItemTag": func() molangValue {
		return molangValue{references: slices.Concat(stores.BlockTag.Source.Get(), stores.ItemTag.Source.Get()), strings: vanilla.ItemTag.ToSlice()}
	},
	"EntityIdentifier": func() molangValue {
		return molangValue{references: stores.EntityId.Source.Get(), strings: vanilla.EntityId.ToSlice()}
	},
	"EntityProperty": func() molangValue {
		return molangValue{references: stores.EntityProperty.Source.Get()}
	},
	"TypeFamily": func() molangValue {
		return molangValue{references: stores.EntityFamily.Source.Get(), strings: vanilla.Family.ToSlice()}
	},
	"ItemIdentifier": func() molangValue {
		return molangValue{references: stores.ItemId.Source.Get(), strings: vanilla.ItemId.ToSlice()}
	},
	"ItemTag": func() molangValue {
		return molangValue{references: stores.ItemTag.Source.Get(), strings: vanilla.ItemTag.ToSlice()}
	},
}
