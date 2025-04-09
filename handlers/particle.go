package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Particle = newJsonHandler(shared.ParticleGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonValue("particle_effect/description/identifier")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("particle_id"), stores.ClientEntity.Get("particle_id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Particle.Get("id")
		},
	},
	{
		Path:       []shared.JsonPath{shared.JsonValue("particle_effect/description/basic_render_parameters/texture")},
		Actions:    completions | definitions,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Texture.GetPaths()
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
})
