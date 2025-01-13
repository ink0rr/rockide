package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var ClientEntity = newJsonHandler(core.ClientEntityGlob, []jsonHandlerEntry{
	{
		Path:       []string{"minecraft:client_entity/description/identifier"},
		MatchType:  "value",
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Entity.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.Get("id")
		},
	},
	{
		Path:       []string{"minecraft:client_entity/description/animations/*"},
		Actions:    completions | definitions | rename,
		MatchType:  "key",
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate")
		},
	},
	{
		Path:      []string{"minecraft:client_entity/description/animations/*"},
		Actions:   completions | definitions | rename,
		MatchType: "value",
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.ClientAnimationController.Get("id"), stores.ClientAnimation.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.Get("animation_id")
		},
	},
	{
		Path:      []string{"minecraft:client_entity/description/scripts/animate/*/*"},
		MatchType: "key",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:      []string{"minecraft:client_entity/description/scripts/animate/*"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:      []string{"minecraft:client_entity/description/textures/*"},
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
		Path:      []string{"minecraft:client_entity/description/geometry/*"},
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
			"minecraft:client_entity/description/render_controllers/*/*",
		},
		MatchType: "key",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.RenderController.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("geometry_id"), stores.ClientEntity.Get("geometry_id"))
		},
	},
	{
		Path: []string{
			"minecraft:client_entity/description/render_controllers/*",
		},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.RenderController.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("geometry_id"), stores.ClientEntity.Get("geometry_id"))
		},
	},
	{
		Path:      []string{"minecraft:client_entity/description/spawn_egg/texture"},
		MatchType: "value",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ItemTexture.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.ClientEntity.Get("spawn_egg"), stores.Item.Get("icon"))
		},
	},
	{
		Path: []string{
			"minecraft:client_entity/description/particle_effects/*",
			"minecraft:client_entity/description/particle_emitters/*",
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
		Path:      []string{"minecraft:client_entity/description/sound_effects/*"},
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
