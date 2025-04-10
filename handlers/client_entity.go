package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientEntity = newJsonHandler(shared.ClientEntityGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/identifier")},
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
		Path:       []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/animations/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate")
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/animations/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.ClientAnimationController.Get("id"), stores.ClientAnimation.Get("id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("animation_id"), stores.ClientEntity.Get("animation_id"))
		},
	},
	{
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:client_entity/description/scripts/animate/*"),
			shared.JsonKey("minecraft:client_entity/description/scripts/animate/*/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientEntity.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/textures/*")},
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Texture.GetPaths()
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/geometry/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Geometry.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("geometry_id"), stores.ClientEntity.Get("geometry_id"))
		},
	},
	{
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:client_entity/description/render_controllers/*"),
			shared.JsonKey("minecraft:client_entity/description/render_controllers/*/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.RenderController.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("render_controller_id"), stores.ClientEntity.Get("render_controller_id"))
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/spawn_egg/texture")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ItemTexture.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.ClientEntity.Get("spawn_egg"), stores.Item.Get("icon"))
		},
	},
	{
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:client_entity/description/particle_effects/*"),
			shared.JsonValue("minecraft:client_entity/description/particle_emitters/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Particle.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("particle_id"), stores.ClientEntity.Get("particle_id"), stores.Particle.Get("particle_id"))
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/sound_effects/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.SoundDefinition.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("sound_definition_id"), stores.ClientEntity.Get("sound_definition_id"))
		},
	},
})
