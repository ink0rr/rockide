package shared

import (
	"slices"

	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/sliceutil"
)

var molangLocations = slices.Concat(
	[]JsonPath{
		JsonValue("animation_controllers/*/states/*/transitions/*/*"),
		JsonValue("animation_controllers/*/states/*/on_entry/*"),
		JsonValue("animation_controllers/*/states/*/on_exit/*"),
		JsonValue("minecraft:biome/components/minecraft:forced_features/*/*/iterations"),
		JsonValue("minecraft:entity/events/**/set_property/*"),
		JsonValue("minecraft:entity/description/scripts/animate/*/*"),
		JsonValue("minecraft:block/permutations/*/condition"),
		JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/repair_amount"),
		JsonValue("minecraft:item/components/minecraft:icon/frame"),
		JsonValue("minecraft:item/components/**/condition"),
		JsonValue("minecraft:item/events/**/sequence/*/condition"),
		JsonValue("minecraft:growing_plant_feature/height_distribution/*/*"),
		JsonValue("minecraft:growing_plant_feature/age"),
		JsonValue("animations/*/anim_time_update"),
		JsonValue("animations/*/bones/*/rotation/*"),
		JsonValue("animations/*/bones/*/rotation/*/*"),
		JsonValue("animations/*/bones/*/scale"),
		JsonValue("animations/*/bones/*/scale/*"),
		JsonValue("animations/*/bones/*/scale/*/*"),
		JsonValue("animations/*/bones/*/position/*"),
		JsonValue("animations/*/bones/*/position/*/*"),
		JsonValue("animations/*/timeline/*"),
		JsonValue("animation_controllers/*/states/*/animations/*/*"),
		JsonValue("animation_controllers/*/states/*/transitions/*/*"),
		JsonValue("animation_controllers/*/states/*/on_entry/*"),
		JsonValue("animation_controllers/*/states/*/on_exit/*"),
		JsonValue("animation_controllers/*/states/*/parameters/*"),
		JsonValue("animation_controllers/*/states/*/variables/*/input"),
		JsonValue("render_controllers/*/uv_anim/offset/*"),
		JsonValue("render_controllers/*/uv_anim/scale/*"),
		JsonValue("render_controllers/*/geometry"),
		JsonValue("render_controllers/*/part_visibility/*/*"),
		JsonValue("render_controllers/*/materials/*/*"),
		JsonValue("render_controllers/*/textures/*"),
		JsonValue("render_controllers/*/color/*"),
		JsonValue("render_controllers/*/overlay_color/*"),
		JsonValue("render_controllers/*/is_hurt_color/*"),
		JsonValue("render_controllers/*/on_fire_color/*"),
		JsonValue("minecraft:attachable/description/scripts/animate/*/*"),
		JsonValue("minecraft:attachable/description/scripts/initialize/*"),
		JsonValue("minecraft:attachable/description/scripts/parent_setup"),
		JsonValue("minecraft:attachable/description/scripts/pre_animation/*"),
		JsonValue("minecraft:attachable/description/scripts/scale"),
		JsonValue("minecraft:attachable/description/render_controllers/*/*"),
		JsonValue("minecraft:client_entity/description/scripts/animate/*/*"),
		JsonValue("minecraft:client_entity/description/scripts/initialize/*"),
		JsonValue("minecraft:client_entity/description/scripts/pre_animation/*"),
		JsonValue("minecraft:client_entity/description/scripts/scale"),
		JsonValue("minecraft:client_entity/description/render_controllers/*/*"),
		JsonValue("particle_effect/curves/*/input"),
		JsonValue("particle_effect/curves/*/horizontal_range"),
		JsonValue("particle_effect/events/**/particle_effect/pre_effect_expression"),

		JsonValue("minecraft:geometry/bones/*/binding"),
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
	}, func(value string) []JsonPath {
		return []JsonPath{
			JsonValue("minecraft:entity/components/" + value),
			JsonValue("minecraft:entity/component_groups/*/" + value),
		}
	}),
	sliceutil.FlatMap([]string{
		"minecraft:emitter_initialization/creation_expression",
		"minecraft:emitter_initialization/per_update_expression",
		"minecraft:emitter_rate_instant/num_particles",
		"minecraft:emitter_rate_steady/spawn_rate",
		"minecraft:emitter_rate_steady/max_particles",
		"minecraft:emitter_rate_manual/max_particles",
		"minecraft:emitter_lifetime_looping/active_time",
		"minecraft:emitter_lifetime_looping/sleep_time",
		"minecraft:emitter_lifetime_once/active_time",
		"minecraft:emitter_lifetime_expression/activation_expression",
		"minecraft:emitter_lifetime_expression/expiration_expression",
		"minecraft:emitter_shape_point/offset/*",
		"minecraft:emitter_shape_point/direction/*",
		"minecraft:emitter_shape_sphere/radius",
		"minecraft:emitter_shape_sphere/offset/*",
		"minecraft:emitter_shape_sphere/direction/*",
		"minecraft:emitter_shape_box/half_dimensions/*",
		"minecraft:emitter_shape_box/offset/*",
		"minecraft:emitter_shape_box/direction/*",
		"minecraft:emitter_shape_custom/offset/*",
		"minecraft:emitter_shape_custom/direction/*",
		"minecraft:emitter_shape_entity_aabb/direction/*",
		"minecraft:emitter_shape_disc/radius",
		"minecraft:emitter_shape_disc/offset/*",
		"minecraft:emitter_shape_disc/direction/*",
		"minecraft:particle_initial_spin/rotation",
		"minecraft:particle_initial_spin/rotation_rate",
		"minecraft:particle_initial_speed",
		"minecraft:particle_initial_speed/*",
		"minecraft:particle_motion_dynamic/linear_acceleration/*",
		"minecraft:particle_motion_dynamic/linear_drag_coefficient",
		"minecraft:particle_motion_dynamic/rotating_acceleration",
		"minecraft:particle_motion_dynamic/rotation_drag_coefficient",
		"minecraft:particle_motion_parametric/relative_position/*",
		"minecraft:particle_motion_parametric/direction/*",
		"minecraft:particle_motion_parametric/rotation",
		"minecraft:particle_motion_collision/enabled",
		"minecraft:particle_appearance_billboard/size/*",
		"minecraft:particle_appearance_billboard/direction/custom_direction/*",
		"minecraft:particle_appearance_billboard/uv/uv/*",
		"minecraft:particle_appearance_billboard/uv/uv_size/*",
		"minecraft:particle_appearance_billboard/uv/flipbook/base_UV/*",
		"minecraft:particle_appearance_billboard/uv/flipbook/max_frame",
		"minecraft:particle_appearance_tinting/color/gradient/*/*",
		"minecraft:particle_appearance/color/interpolant",
		"minecraft:particle_lifetime_expression/expiration_expression",
		"minecraft:particle_lifetime_expression/max_lifetime",
		"minecraft:particle_initialization/per_update_expression",
		"minecraft:particle_initialization/per_render_expression",
	}, func(value string) []JsonPath {
		return []JsonPath{
			JsonValue("particle_effect/components/" + value),
			JsonValue("particle_effect/events/**/components/" + value),
		}
	}),
)

// For locations that uses the highlighting but not the completions
var MolangSemanticLocations = slices.Concat(molangLocations, []JsonPath{
	JsonValue("render_controllers/*/arrays/*/*/*"),
	JsonValue("minecraft:attachable/description/geometry/*"),
	JsonValue("minecraft:client_entity/description/geometry/*"),
	JsonValue("minecraft:geometry/*/description/identifier"),
})

func IsMolangLocation(location *jsonc.Location) bool {
	node := location.PreviousNode
	if node == nil {
		return false
	}
	nodeValue, ok := node.Value.(string)
	if !ok || nodeValue == "" {
		return false
	}
	return node.Type == jsonc.NodeTypeString &&
		nodeValue[0] != '@' &&
		nodeValue[0] != '/' &&
		slices.ContainsFunc(molangLocations, func(jsonPath JsonPath) bool {
			return jsonPath.IsKey == location.IsAtPropertyKey && location.Path.Matches(jsonPath.Path)
		})
}
