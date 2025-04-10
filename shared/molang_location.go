package shared

import (
	"slices"

	"github.com/ink0rr/rockide/internal/jsonc"
)

var molangLocations = []JsonPath{
	JsonValue("animation_controllers/*/states/*/transitions/*/*"),
	JsonValue("animation_controllers/*/states/*/on_entry/*"),
	JsonValue("animation_controllers/*/states/*/on_exit/*"),
	JsonValue("minecraft:biome/components/minecraft:forced_features/*/*/iterations"),
	JsonValue("minecraft:entity/events/**/set_property/*"),
	JsonValue("minecraft:entity/description/scripts/animate/*/*"),
	JsonValue("minecraft:entity/components/minecraft:behavior.eat_block/success_chance"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:behavior.eat_block/success_chance"),
	JsonValue("minecraft:entity/components/minecraft:experience_reward/on_bred"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:experience_reward/on_bred"),
	JsonValue("minecraft:entity/components/minecraft:experience_reward/on_death"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:experience_reward/on_death"),
	JsonValue("minecraft:entity/components/minecraft:projectile/on_hit/impact_damage/filter"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:projectile/on_hit/impact_damage/filter"),
	JsonValue("minecraft:entity/components/minecraft:rideable/seats/*/rotate_rider_by"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:rideable/seats/*/rotate_rider_by"),
	JsonValue("minecraft:entity/components/minecraft:rideable/seats/rotate_rider_by"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:rideable/seats/rotate_rider_by"),
	JsonValue("minecraft:entity/components/minecraft:ambient_sound_interval/event_names/*/condition"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:ambient_sound_interval/event_names/*/condition"),
	JsonValue("minecraft:entity/components/minecraft:anger_level/on_increase_sounds/*/condition"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:anger_level/on_increase_sounds/*/condition"),
	JsonValue("minecraft:entity/components/minecraft:heartbeat/interval"),
	JsonValue("minecraft:entity/component_groups/*/minecraft:heartbeat/interval"),
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
	JsonValue("particle_effect/events/*/particle_effect/pre_effect_expression"),
	JsonValue("particle_effect/components/minecraft:emitter_initialization/creation_expression"),
	JsonValue("particle_effect/components/minecraft:emitter_initializtion/per_update_expression"),
	JsonValue("particle_effect/components/minecraft:emitter_rate_instant/num_particles"),
	JsonValue("particle_effect/components/minecraft:emitter_rate_steady/spawn_rate"),
	JsonValue("particle_effect/components/minecraft:emitter_rate_steady/max_particles"),
	JsonValue("particle_effect/components/minecraft:emitter_rate_manual/max_particles"),
	JsonValue("particle_effect/components/minecraft:emitter_lifetime_looping/active_time"),
	JsonValue("particle_effect/components/minecraft:emitter_lifetime_looping/sleep_time"),
	JsonValue("particle_effect/components/minecraft:emitter_lifetime_once/active_time"),
	JsonValue("particle_effect/components/minecraft:emitter_lifetime_expression/activation_expression"),
	JsonValue("particle_effect/components/minecraft:emitter_lifetime_expression/expiration_expression"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_point/offset/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_point/direction/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_sphere/radius"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_sphere/offset/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_sphere/direction/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_box/half_dimensions/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_box/offset/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_box/direction/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_custom/offset/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_custom/direction/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_entity_aabb/direction/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_disc/radius"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_disc/offset/*"),
	JsonValue("particle_effect/components/minecraft:emitter_shape_disc/direction/*"),
	JsonValue("particle_effect/components/minecraft:particle_initial_spin/rotation"),
	JsonValue("particle_effect/components/minecraft:particle_initial_spin/rotation_rate"),
	JsonValue("particle_effect/components/minecraft:particle_initial_speed"),
	JsonValue("particle_effect/components/minecraft:particle_initial_speed/*"),
	JsonValue("particle_effect/components/minecraft:particle_motion_dynamic/linear_acceleration/*"),
	JsonValue("particle_effect/components/minecraft:particle_motion_dynamic/linear_drag_coefficient"),
	JsonValue("particle_effect/components/minecraft:particle_motion_dynamic/rotating_acceleration"),
	JsonValue("particle_effect/components/minecraft:particle_motion_dynamic/rotation_drag_coefficient"),
	JsonValue("particle_effect/components/minecraft:particle_motion_parametric/relative_position/*"),
	JsonValue("particle_effect/components/minecraft:particle_motion_parametric/direction/*"),
	JsonValue("particle_effect/components/minecraft:particle_motion_parametric/rotation"),
	JsonValue("particle_effect/components/minecraft:particle_motion_collision/enabled"),
	JsonValue("particle_effect/components/minecraft:particle_appearance_billboard/size/*"),
	JsonValue("particle_effect/components/minecraft:particle_appearance_billboard/direction/custom_direction/*"),
	JsonValue("particle_effect/components/minecraft:particle_appearance_billboard/uv/uv/*"),
	JsonValue("particle_effect/components/minecraft:particle_appearance_billboard/uv/uv_size/*"),
	JsonValue("particle_effect/components/minecraft:particle_appearance_billboard/uv/flipbook/base_UV/*"),
	JsonValue("particle_effect/components/minecraft:particle_appearance_billboard/uv/flipbook/max_frame"),
	JsonValue("particle_effect/components/minecraft:particle_appearance_tinting/color/gradient/*/*"),
	JsonValue("particle_effect/components/minecraft:particle_appearance/color/interpolant"),
	JsonValue("particle_effect/components/minecraft:particle_lifetime_expression/expiration_expression"),
	JsonValue("particle_effect/components/minecraft:particle_lifetime_expression/max_lifetime"),
	JsonValue("particle_effect/components/minecraft:particle_initialization/per_update_expression"),
	JsonValue("particle_effect/components/minecraft:particle_initialization/per_render_expression"),
	JsonValue("minecraft:geometry/bones/*/binding"),
}

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
	if !ok {
		return false
	}
	return node.Type == jsonc.NodeTypeString &&
		nodeValue[0] != '@' &&
		nodeValue[0] != '/' &&
		slices.ContainsFunc(molangLocations, func(jsonPath JsonPath) bool {
			return jsonPath.IsKey == location.IsAtPropertyKey && location.Path.Matches(jsonPath.Path)
		})
}
