package stores

import "github.com/ink0rr/rockide/shared"

var Particle = &JsonStore{
	pattern: shared.ParticleGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("particle_effect/description/identifier")},
		},
		{
			Id:   "id_refs",
			Path: []shared.JsonPath{shared.JsonValue("particle_effect/events/**/particle_effect/effect")},
		},
		{
			Id:   "texture_path",
			Path: []shared.JsonPath{shared.JsonValue("particle_effect/description/basic_render_parameters/texture")},
		},
		{
			Id:   "event",
			Path: []shared.JsonPath{shared.JsonKey("particle_effect/events/*")},
		},
		{
			Id: "event_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("particle_effect/components/minecraft:emitter_lifetime_events/creation_event"),
				shared.JsonValue("particle_effect/components/minecraft:emitter_lifetime_events/expiration_event"),
				shared.JsonValue("particle_effect/components/minecraft:emitter_lifetime_events/looping_travel_distance_events"),
				shared.JsonValue("particle_effect/components/minecraft:emitter_lifetime_events/timeline/*/*"),
				shared.JsonValue("particle_effect/components/minecraft:emitter_lifetime_events/travel_distance_events/*/*"),
				shared.JsonValue("particle_effect/components/minecraft:particle_lifetime_events/creation_event"),
				shared.JsonValue("particle_effect/components/minecraft:particle_lifetime_events/expiration_event"),
				shared.JsonValue("particle_effect/components/minecraft:particle_lifetime_events/timeline/*/*"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:emitter_lifetime_events/creation_event"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:emitter_lifetime_events/expiration_event"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:emitter_lifetime_events/looping_travel_distance_events"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:emitter_lifetime_events/timeline/*/*"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:emitter_lifetime_events/travel_distance_events/*/*"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:particle_lifetime_events/creation_event"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:particle_lifetime_events/expiration_event"),
				shared.JsonValue("particle_effect/events/**/components/minecraft:particle_lifetime_events/timeline/*/*"),
			},
		},
	},
}
