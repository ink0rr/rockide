package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/sliceutil"
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

var inputModes = [...]string{"keyboard_and_mouse", "touch", "gamepad", "motion_controller"}

var molangLiterals = map[string]func() []string{
	"BiomeTag":      func() []string { return biomeTags[:] },
	"EquipmentSlot": func() []string { return equipmentSlots[:] },
	"InputMode":     func() []string { return inputModes[:] },
	"BlockTag": func() []string {
		return sliceutil.Map(stores.Block.Get("tag"), func(ref core.Reference) string { return ref.Value })
	},
	"BlockAndItemTag": func() []string {
		return sliceutil.Map(slices.Concat(stores.Block.Get("tag"), stores.Item.Get("tag")), func(ref core.Reference) string { return ref.Value })
	},
	"EntityIdentifier": func() []string {
		return sliceutil.Map(stores.Entity.Get("id"), func(ref core.Reference) string { return ref.Value })
	},
	"EntityProperty": func() []string {
		return sliceutil.Map(stores.Entity.Get("property"), func(ref core.Reference) string { return ref.Value })
	},
	"TypeFamily": func() []string {
		return sliceutil.Map(stores.Entity.Get("family"), func(ref core.Reference) string { return ref.Value })
	},
	"ItemIdentifier": func() []string {
		return sliceutil.Map(stores.Item.Get("id"), func(ref core.Reference) string { return ref.Value })
	},
	"ItemTag": func() []string {
		return sliceutil.Map(stores.Item.Get("tag"), func(ref core.Reference) string { return ref.Value })
	},
}
