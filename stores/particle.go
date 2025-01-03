package stores

import "github.com/ink0rr/rockide/core"

var Particle = newJsonStore(core.ParticleGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"particle_effect/description/identifier"},
	},
	{
		Id:   "identifier_refs",
		Path: []string{"particle_effect/events/**/particle_effect/effect"},
	},
	{
		Id:   "texture_path",
		Path: []string{"particle_effect/description/basic_render_parameters/texture"},
	},
	{
		Id:   "event",
		Path: []string{"particle_effect/events"},
	},
	{
		Id: "event_refs",
		Path: []string{
			"particle_effect/components/minecraft:emitter_lifetime_events/creation_event",
			"particle_effect/components/minecraft:emitter_lifetime_events/expiration_event",
			"particle_effect/components/minecraft:emitter_lifetime_events/looping_travel_distance_events",
			"particle_effect/components/minecraft:emitter_lifetime_events/timeline/*/*",
			"particle_effect/components/minecraft:emitter_lifetime_events/travel_distance_events/*/*",
			"particle_effect/components/minecraft:particle_lifetime_events/creation_event",
			"particle_effect/components/minecraft:particle_lifetime_events/expiration_event",
			"particle_effect/components/minecraft:particle_lifetime_events/timeline/*/*",
			"particle_effect/events/**/components/minecraft:emitter_lifetime_events/creation_event",
			"particle_effect/events/**/components/minecraft:emitter_lifetime_events/expiration_event",
			"particle_effect/events/**/components/minecraft:emitter_lifetime_events/looping_travel_distance_events",
			"particle_effect/events/**/components/minecraft:emitter_lifetime_events/timeline/*/*",
			"particle_effect/events/**/components/minecraft:emitter_lifetime_events/travel_distance_events/*/*",
			"particle_effect/events/**/components/minecraft:particle_lifetime_events/creation_event",
			"particle_effect/events/**/components/minecraft:particle_lifetime_events/expiration_event",
			"particle_effect/events/**/components/minecraft:particle_lifetime_events/timeline/*/*",
		},
	},
})
