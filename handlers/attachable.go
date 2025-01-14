package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var Attachable = newJsonHandler(core.AttachableGlob, []jsonHandlerEntry{
	{
		Path:       []string{"minecraft:attachable/description/identifier"},
		MatchType:  "value",
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Attachable.Get("id")
		},
	},
	{
		Path:       []string{"minecraft:attachable/description/animations/*"},
		Actions:    completions | definitions | rename,
		MatchType:  "key",
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate")
		},
	},
	{
		Path:      []string{"minecraft:attachable/description/animations/*"},
		Actions:   completions | definitions | rename,
		MatchType: "value",
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.ClientAnimationController.Get("id"), stores.ClientAnimation.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("animation_id"), stores.ClientEntity.Get("animation_id"))
		},
	},
	{
		Path:      []string{"minecraft:attachable/description/scripts/animate/*/*"},
		MatchType: "key",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:      []string{"minecraft:attachable/description/scripts/animate/*"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:      []string{"minecraft:attachable/description/textures/*"},
		MatchType: "value",
		Actions:   completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Texture.Get("path")
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Path:      []string{"minecraft:attachable/description/geometry/*"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Geometry.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("geometry_id"), stores.ClientEntity.Get("geometry_id"))
		},
	},
	{
		Path: []string{
			"minecraft:attachable/description/render_controllers/*/*",
		},
		MatchType: "key",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.RenderController.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("render_controller_id"), stores.ClientEntity.Get("render_controller_id"))
		},
	},
	{
		Path: []string{
			"minecraft:attachable/description/render_controllers/*",
		},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.RenderController.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("render_controller_id"), stores.ClientEntity.Get("render_controller_id"))
		},
	},
	{
		Path: []string{
			"minecraft:attachable/description/particle_effects/*",
			"minecraft:attachable/description/particle_emitters/*",
		},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Particle.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Particle.Get("id_refs"), stores.Attachable.Get("particle_id"), stores.ClientEntity.Get("particle_id"))
		},
	},
	{
		Path:      []string{"minecraft:attachable/description/sound_effects/*"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.SoundDefinition.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("sound_definition_id"), stores.ClientEntity.Get("sound_definition_id"))
		},
	},
})
