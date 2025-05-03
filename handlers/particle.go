package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var Particle = &JsonHandler{Pattern: shared.ParticleGlob}

func init() {
	Particle.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonValue("particle_effect/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("particle_id"), ClientEntity.Get("particle_id"), Particle.Get("particle_id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Particle.Get("id")
			},
		},
		{
			Id:   "particle_id",
			Path: []shared.JsonPath{shared.JsonValue("particle_effect/events/**/particle_effect/effect")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Particle.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("particle_id"), ClientEntity.Get("particle_id"), Particle.Get("particle_id"))
			},
			VanillaData: vanilla.ParticleIdentifiers,
		},
		{
			Id:            "texture_path",
			Path:          []shared.JsonPath{shared.JsonValue("particle_effect/description/basic_render_parameters/texture")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Texture.GetPaths()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.TexturePaths,
		},
		{
			Id:         "event",
			Path:       []shared.JsonPath{shared.JsonKey("particle_effect/events/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Particle.GetFrom(ctx.URI, "event_refs")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Particle.GetFrom(ctx.URI, "event")
			},
		},
		{
			Id: "event_refs",
			Path: sliceutil.FlatMap([]string{
				"minecraft:emitter_lifetime_events/creation_event",
				"minecraft:emitter_lifetime_events/expiration_event",
				"minecraft:emitter_lifetime_events/looping_travel_distance_events/*/effects",
				"minecraft:emitter_lifetime_events/timeline/*",
				"minecraft:emitter_lifetime_events/travel_distance_events/*",
				"minecraft:particle_lifetime_events/creation_event",
				"minecraft:particle_lifetime_events/expiration_event",
				"minecraft:particle_lifetime_events/timeline/*",
			}, func(path string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("particle_effect/components/" + path),
					shared.JsonValue("particle_effect/components/" + path + "/*"),
					shared.JsonValue("particle_effect/events/**/components/" + path),
					shared.JsonValue("particle_effect/events/**/components/" + path + "/*"),
				}
			}),
			Source: func(ctx *JsonContext) []core.Symbol {
				return Particle.GetFrom(ctx.URI, "event")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Particle.GetFrom(ctx.URI, "event_refs")
			},
		},
	}
	Particle.MolangLocations = slices.Concat(
		[]shared.JsonPath{
			shared.JsonValue("particle_effect/curves/*/input"),
			shared.JsonValue("particle_effect/curves/*/horizontal_range"),
			shared.JsonValue("particle_effect/events/**/particle_effect/pre_effect_expression"),
		},
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
		}, func(value string) []shared.JsonPath {
			return []shared.JsonPath{
				shared.JsonValue("particle_effect/components/" + value),
				shared.JsonValue("particle_effect/events/**/components/" + value),
			}
		}),
	)
}
