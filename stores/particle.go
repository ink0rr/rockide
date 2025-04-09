package stores

import (
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
)

var Particle = &JsonStore{
	pattern: shared.ParticleGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("particle_effect/description/identifier")},
		},
		{
			Id:   "particle_id",
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
		},
	},
}
