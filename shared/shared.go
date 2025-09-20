package shared

import (
	"os"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/sliceutil"
)

var wd string

func Getwd() string {
	if wd == "" {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		wd = dir
	}
	return wd
}

var project *core.Project

func GetProject() *core.Project {
	return project
}

func SetProject(p *core.Project) {
	project = p
}

const (
	BpGlob      = "{behavior_pack,*BP,BP_*,*bp,bp_*}"
	RpGlob      = "{resource_pack,*RP,RP_*,*rp,rp_*}"
	ProjectGlob = "{behavior_pack,*BP,BP_*,*bp,bp_*,resource_pack,*RP,RP_*,*rp,rp_*}"
)

type Pattern string

func (p Pattern) PackType() string {
	switch p[0] {
	case 'b':
		return project.BP
	case 'r':
		return project.RP
	default:
		panic("invalid pattern")
	}
}

func (p Pattern) ToString() string {
	return p.PackType() + string(p[1:])
}

func behaviorPattern(pattern string) Pattern {
	return Pattern("b" + pattern)
}

func resourcePattern(pattern string) Pattern {
	return Pattern("r" + pattern)
}

var (
	AimAssistPresetsGlob    = behaviorPattern("/aim_assist/presets/**/*.json")
	AimAssistCategoriesGlob = behaviorPattern("/aim_assist/categories/**/*.json")
	AnimationControllerGlob = behaviorPattern("/animation_controllers/**/*.json")
	AnimationGlob           = behaviorPattern("/animations/**/*.json")
	BiomeGlob               = behaviorPattern("/biomes/**/*.json")
	BlockGlob               = behaviorPattern("/blocks/**/*.json")
	CameraGlob              = behaviorPattern("/cameras/presets/**/*.json")
	CraftingItemCatalogGlob = behaviorPattern("/item_catalog/crafting_item_catalog.json")
	DialogueGlob            = behaviorPattern("/dialogue/**/*.json")
	EntityGlob              = behaviorPattern("/entities/**/*.json")
	FeatureRuleGlob         = behaviorPattern("/feature_rules/**/*.json")
	FeatureGlob             = behaviorPattern("/features/**/*.json")
	FunctionGlob            = behaviorPattern("/functions/**/*.mcfunction")
	ItemGlob                = behaviorPattern("/items/**/*.json")
	LootTableGlob           = behaviorPattern("/loot_tables/**/*.json")
	RecipeGlob              = behaviorPattern("/recipes/**/*.json")
	SpawnRuleGlob           = behaviorPattern("/spawn_rules/**/*.json")
	StructureGlob           = behaviorPattern("/structures/**/*.mcstructure")
	TradeTableGlob          = behaviorPattern("/trading/**/*.json")
)

var (
	AttachableGlob                = resourcePattern("/attachables/**/*.json")
	AtmosphereGlob                = resourcePattern("/atmospherics/**/*.json")
	ClientAnimationControllerGlob = resourcePattern("/animation_controllers/**/*.json")
	ClientAnimationGlob           = resourcePattern("/animations/**/*.json")
	ClientBiomeGlob               = resourcePattern("/biomes/**/*.json")
	ClientBlockGlob               = resourcePattern("/blocks.json")
	ClientEntityGlob              = resourcePattern("/entity/**/*.json")
	ClientSoundGlob               = resourcePattern("/sounds.json")
	ColorGradingGlob              = resourcePattern("/color_grading/**/*.json")
	FogGlob                       = resourcePattern("/fogs/**/*.json")
	GeometryGlob                  = resourcePattern("/models/**/*.json")
	ItemTextureGlob               = resourcePattern("/textures/item_texture.json")
	LightingGlob                  = resourcePattern("/lighting/**/*.json")
	MusicDefinitionGlob           = resourcePattern("/sounds/music_definitions.json")
	ParticleGlob                  = resourcePattern("/particles/**/*.json")
	RenderControllerGlob          = resourcePattern("/render_controllers/**/*.json")
	SoundDefinitionGlob           = resourcePattern("/sounds/sound_definitions.json")
	SoundGlob                     = resourcePattern("/sounds/**/*.{fsb,ogg,wav}")
	TerrainTextureGlob            = resourcePattern("/textures/terrain_texture.json")
	TextureGlob                   = resourcePattern("/textures/**/*.{png,tga,fsb}")
	WaterGlob                     = resourcePattern("/water/**/*.json")
)

var PropertyTests = []string{"bool_property", "enum_property", "float_property", "int_property"}

var FilterPaths = sliceutil.FlatMap([]string{
	"**/filters",
	"minecraft:ageable/interact_filters",
	"minecraft:anger_level/nuisance_filter",
	"minecraft:angry/broadcast_filters",
	"minecraft:area_attack/entity_filter",
	"minecraft:behavior.knockback_roar/damage_filters",
	"minecraft:behavior.knockback_roar/knockback_filters",
	"minecraft:behavior.move_to_block/target_block_filters",
	"minecraft:behavior.nap/can_nap_filters",
	"minecraft:behavior.nap/wake_mob_exceptions",
	"minecraft:behavior.stalk_and_pounce_on_target/stuck_blocks",
	"minecraft:block_sensor/sources",
	"minecraft:breedable/love_filters",
	"minecraft:celebrate_hunt/celebration_targets",
	"minecraft:conditional_bandwidth_optimization/conditional_values",
	"minecraft:entity_sensor/event_filters",
	"minecraft:entity_sensor/subsensors/*/event_filters",
	"minecraft:mob_effect/entity_filter",
	"minecraft:trail/spawn_filter",
}, func(path string) []string {
	return []string{
		"minecraft:entity/components/" + path + "/**",
		"minecraft:entity/component_groups/*/" + path + "/**",
		"minecraft:entity/events/*/" + path + "/**",
	}
})
