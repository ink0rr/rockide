package core

const (
	BpGlob = "{behavior_pack,*BP,BP_*,*bp,bp_*}"
	RpGlob = "{resource_pack,*RP,RP_*,*rp,rp_*}"
)

type Pattern string

func (p Pattern) Resolve(project *Project) string {
	pattern := string(p[1:])
	switch p[0] {
	case 'b':
		return project.BP + pattern
	case 'r':
		return project.RP + pattern
	default:
		panic("invalid pattern")
	}
}

func behaviorPattern(pattern string) Pattern {
	return Pattern("b" + pattern)
}

func resourcePattern(pattern string) Pattern {
	return Pattern("r" + pattern)
}

var (
	AnimationControllerGlob = behaviorPattern("/animation_controllers/**/*.json")
	AnimationGlob           = behaviorPattern("/animations/**/*.json")
	BlockGlob               = behaviorPattern("/blocks/**/*.json")
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
	ClientAnimationControllerGlob = resourcePattern("/animation_controllers/**/*.json")
	ClientAnimationGlob           = resourcePattern("/animations/**/*.json")
	ClientBlockGlob               = resourcePattern("/blocks.json")
	ClientEntityGlob              = resourcePattern("/entity/**/*.json")
	GeometryGlob                  = resourcePattern("/models/**/*.json")
	ItemTextureGlob               = resourcePattern("/textures/item_texture.json")
	ParticleGlob                  = resourcePattern("/particles/**/*.json")
	RenderControllerGlob          = resourcePattern("/render_controllers/**/*.json")
	SoundDefinitionGlob           = resourcePattern("/sounds/sound_definitions.json")
	SoundGlob                     = resourcePattern("/sounds/**/*.{fsb,ogg,wav}")
	TerrainTextureGlob            = resourcePattern("/textures/terrain_texture.json")
	TextureGlob                   = resourcePattern("/textures/**/*.{png,tga,fsb}")
)

var PropertyDomain = []string{"bool_property", "enum_property", "float_property", "int_property"}
