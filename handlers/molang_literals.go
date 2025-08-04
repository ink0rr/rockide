package handlers

import (
	"github.com/ink0rr/rockide/stores"
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
	bindings []*stores.SymbolBinding
	literals []string
}

var molangTypes = map[string]func() molangValue{
	"BiomeTag": func() molangValue {
		return molangValue{literals: biomeTags[:]}
	},
	"EquipmentSlot": func() molangValue {
		return molangValue{literals: equipmentSlots[:]}
	},
	"GraphicsMode": func() molangValue {
		return molangValue{literals: graphicsModes[:]}
	},
	"InputMode": func() molangValue {
		return molangValue{literals: inputModes[:]}
	},
	"BlockTag": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.BlockTag}}
	},
	"BlockAndItemTag": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.BlockTag, stores.ItemTag}}
	},
	"EntityIdentifier": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.EntityId}}
	},
	"EntityProperty": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.EntityProperty}}
	},
	"TypeFamily": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.EntityFamily}}
	},
	"ItemIdentifier": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.ItemId}}
	},
	"ItemTag": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.ItemTag}}
	},
}
