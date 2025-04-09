package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Attachable = newJsonHandler(shared.AttachableGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/identifier")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Item.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("id"), stores.Entity.Get("item_id"), stores.Item.Get("item_id"), stores.LootTable.Get("item_id"), stores.Recipe.Get("item_id"))
		},
	},
	{
		Path:       []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/animations/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate_refs")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate")
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/animations/*")},
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
			shared.JsonKey("minecraft:attachable/description/scripts/animate/*/*"),
			shared.JsonValue("minecraft:attachable/description/scripts/animate/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Attachable.GetFrom(params.URI, "animate_refs")
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/textures/*")},
		Actions: completions | definitions,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Texture.GetPaths()
		},
		References: func(params *jsonParams) []core.Reference {
			return nil
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/geometry/*")},
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
			shared.JsonKey("minecraft:attachable/description/render_controllers/*/*"),
			shared.JsonValue("minecraft:attachable/description/render_controllers/*"),
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
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:attachable/description/particle_effects/*"),
			shared.JsonValue("minecraft:attachable/description/particle_emitters/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Particle.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Particle.Get("id_refs"), stores.Attachable.Get("particle_id"), stores.ClientEntity.Get("particle_id"))
		},
	},
	{
		Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/sound_effects/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.SoundDefinition.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Attachable.Get("sound_definition_id"), stores.ClientEntity.Get("sound_definition_id"))
		},
	},
})
