package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var ClientEntity = &JsonHandler{Pattern: shared.ClientEntityGlob}

func init() {
	ClientEntity.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Entity.Get("id"), Entity.Get("id_refs"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ClientEntity.Get("id")
			},
		},
		{
			Id:         "animate",
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return ClientEntity.GetFrom(ctx.URI, "animate_refs")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ClientEntity.GetFrom(ctx.URI, "animate")
			},
		},
		{
			Id:   "animation_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/animations/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(ClientAnimationController.Get("id"), ClientAnimation.Get("id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("animation_id"), ClientEntity.Get("animation_id"))
			},
			VanillaData: vanilla.AnimationAndController,
		},
		{
			Id: "animate_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:client_entity/description/scripts/animate/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return ClientEntity.GetFrom(ctx.URI, "animate")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ClientEntity.GetFrom(ctx.URI, "animate_refs")
			},
		},
		{
			Id:   "material",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/materials/*")},
			// TODO
		},
		{
			Id:   "material_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/materials/*")},
			// TODO
		},
		{
			Id:   "texture",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/textures/*")},
		},
		{
			Id:            "texture_path",
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/textures/*")},
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
			Id:   "geometry",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/geometry/*")},
			// TODO
		},
		{
			Id:   "geometry_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/geometry/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Geometry.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("geometry_id"), Block.Get("geometry_id"), ClientEntity.Get("geometry_id"))
			},
			VanillaData: vanilla.Geometry,
		},
		{
			Id: "render_controller_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/render_controllers/*"),
				shared.JsonKey("minecraft:client_entity/description/render_controllers/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return RenderController.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("render_controller_id"), ClientEntity.Get("render_controller_id"))
			},
			VanillaData: vanilla.RenderController,
		},
		{
			Id:   "spawn_egg",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/spawn_egg/texture")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return ItemTexture.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(ClientEntity.Get("spawn_egg"), Item.Get("icon"))
			},
			VanillaData: vanilla.ItemTexture,
		},
		{
			Id: "particle",
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonKey("minecraft:client_entity/description/particle_emitters/*"),
			},
			// TODO
		},
		{
			Id: "particle_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonValue("minecraft:client_entity/description/particle_emitters/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Particle.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("particle_id"), Particle.Get("particle_id"), ClientEntity.Get("particle_id"))
			},
			VanillaData: vanilla.ParticleIdentifiers,
		},
		{
			Id:   "sound_definition",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/sound_effects/*")},
			// TODO
		},
		{
			Id:   "sound_definition_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/sound_effects/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return SoundDefinition.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("sound_definition_id"), ClientEntity.Get("sound_definition_id"))
			},
			VanillaData: vanilla.SoundDefinition,
		},
	}
	ClientEntity.MolangLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:client_entity/description/scripts/animate/*/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/initialize/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/parent_setup"),
		shared.JsonValue("minecraft:client_entity/description/scripts/pre_animation/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/scale"),
		shared.JsonValue("minecraft:client_entity/description/render_controllers/*/*"),
	}
	ClientEntity.MolangSemanticLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:client_entity/description/geometry/*"),
	}
}
