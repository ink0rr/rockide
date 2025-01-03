package core

const BpGlob = "{behavior_pack,*BP,BP_*,*bp,bp_*}"
const RpGlob = "{resource_pack,*RP,RP_*,*rp,rp_*}"
const ProjectGlob = "{behavior_pack,*BP,BP_*,*bp,bp_*,resource_pack,*RP,RP_*,*rp,rp_*}"

const (
	AnimationControllerGlob = BpGlob + "/animation_controllers/**/*.json"
	AnimationGlob           = BpGlob + "/animations/**/*.json"
	BlockGlob               = BpGlob + "/blocks/**/*.json"
	EntityGlob              = BpGlob + "/entities/**/*.json"
	FeatureRuleGlob         = BpGlob + "/feature_rules/**/*.json"
	FeatureGlob             = BpGlob + "/features/**/*.json"
	FunctionGlob            = BpGlob + "/functions/**/*.mcfunction"
	ItemGlob                = BpGlob + "/items/**/*.json"
	LootTableGlob           = BpGlob + "/loot_tables/**/*.json"
	RecipeGlob              = BpGlob + "/recipes/**/*.json"
	SpawnRuleGlob           = BpGlob + "/spawn_rules/**/*.json"
	StructureGlob           = BpGlob + "/structures/**/*.mcstructure"
	TradeTableGlob          = BpGlob + "/trading/**/*.json"
)

const (
	AttachableGlob                = RpGlob + "/attachables/**/*.json"
	ClientAnimationControllerGlob = RpGlob + "/animation_controllers/**/*.json"
	ClientAnimationGlob           = RpGlob + "/animations/**/*.json"
	ClientBlockGlob               = RpGlob + "/blocks.json"
	ClientEntityGlob              = RpGlob + "/entity/**/*.json"
	GeometryGlob                  = RpGlob + "/models/**/*.json"
	ItemTextureGlob               = RpGlob + "/textures/item_texture.json"
	ParticleGlob                  = RpGlob + "/particles/**/*.json"
	RenderControllerGlob          = RpGlob + "/render_controllers/**/*.json"
	SoundDefinitionGlob           = RpGlob + "/sounds/sound_definitions.json"
	SoundGlob                     = RpGlob + "/sounds/**/*.{fsb,ogg,wav}"
	TerrainTextureGlob            = RpGlob + "/textures/terrain_texture.json"
	TextureGlob                   = RpGlob + "/textures/**/*.{png,tga,fsb}"
)

var PropertyDomain = []string{"bool_property", "enum_property", "float_property", "int_property"}
