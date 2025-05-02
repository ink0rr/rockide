package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/vanilla"
)

var Particle = &jsonHandler{
	pattern: shared.ParticleGlob,
	entries: []jsonHandlerEntry{
		{
			Path:       []shared.JsonPath{shared.JsonValue("particle_effect/description/identifier")},
			Actions:    completions | definitions | rename,
			FilterDiff: true,
			Source: func(params *jsonParams) []core.Reference {
				return slices.Concat(stores.Attachable.Get("particle_id"), stores.ClientEntity.Get("particle_id"), stores.Particle.Get("particle_id"))
			},
			References: func(params *jsonParams) []core.Reference {
				return stores.Particle.Get("id")
			},
		},
		{
			Path:    []shared.JsonPath{shared.JsonValue("particle_effect/events/**/particle_effect/effect")},
			Actions: completions | definitions | rename,
			Source: func(params *jsonParams) []core.Reference {
				return stores.Particle.Get("id")
			},
			References: func(params *jsonParams) []core.Reference {
				return slices.Concat(stores.Attachable.Get("particle_id"), stores.ClientEntity.Get("particle_id"), stores.Particle.Get("particle_id"))
			},
			VanillaData: vanilla.ParticleIdentifiers,
		},
		{
			Path:    []shared.JsonPath{shared.JsonValue("particle_effect/description/basic_render_parameters/texture")},
			Actions: completions | definitions,
			Source: func(params *jsonParams) []core.Reference {
				return stores.Texture.GetPaths()
			},
			References: func(params *jsonParams) []core.Reference {
				return nil
			},
			VanillaData: vanilla.TexturePaths,
		},
		{
			Path:       []shared.JsonPath{shared.JsonKey("particle_effect/events/*")},
			Actions:    completions | definitions | rename,
			FilterDiff: true,
			Source: func(params *jsonParams) []core.Reference {
				return stores.Particle.GetFrom(params.URI, "event_refs")
			},
			References: func(params *jsonParams) []core.Reference {
				return stores.Particle.GetFrom(params.URI, "event")
			},
		},
		{
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
			Actions: completions | definitions | rename,
			Source: func(params *jsonParams) []core.Reference {
				return stores.Particle.GetFrom(params.URI, "event")
			},
			References: func(params *jsonParams) []core.Reference {
				return stores.Particle.GetFrom(params.URI, "event_refs")
			},
		},
	},
}
