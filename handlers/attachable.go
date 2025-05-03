package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var Attachable = &JsonHandler{Pattern: shared.AttachableGlob}

func init() {
	Attachable.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Item.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("id"), Entity.Get("item_id"), Item.Get("item_id"), LootTable.Get("item_id"), Recipe.Get("item_id"), TradeTable.Get("item_id"))
			},
		},
		{
			Id:         "animate",
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Attachable.GetFrom(ctx.URI, "animate_refs")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Attachable.GetFrom(ctx.URI, "animate")
			},
		},
		{
			Id:   "animation_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/animations/*")},
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
				shared.JsonKey("minecraft:attachable/description/scripts/animate/*/*"),
				shared.JsonValue("minecraft:attachable/description/scripts/animate/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Attachable.GetFrom(ctx.URI, "animate")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Attachable.GetFrom(ctx.URI, "animate_refs")
			},
		},
		{
			Id:   "material",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/materials/*")},
			// TODO
		},
		{
			Id:   "material_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/materials/*")},
			// TODO
		},
		{
			Id:   "texture",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/textures/*")},
			// TODO
		},
		{
			Id:            "texture_path",
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/textures/*")},
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
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/geometry/*")},
			// TODO
		},
		{
			Id:   "geometry_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/geometry/*")},
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
				shared.JsonKey("minecraft:attachable/description/render_controllers/*/*"),
				shared.JsonValue("minecraft:attachable/description/render_controllers/*"),
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
			Id: "particle",
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/particle_effects/*"),
				shared.JsonKey("minecraft:attachable/description/particle_emitters/*"),
			},
		},
		{
			Id: "particle_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:attachable/description/particle_effects/*"),
				shared.JsonValue("minecraft:attachable/description/particle_emitters/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Particle.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("particle_id"), ClientEntity.Get("particle_id"), Particle.Get("particle_id"))
			},
			VanillaData: vanilla.ParticleIdentifiers,
		},
		{
			Id:   "sound_definition",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/sound_effects/*")},
			// TODO
		},
		{
			Id:   "sound_definition_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/sound_effects/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return SoundDefinition.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("sound_definition_id"), ClientEntity.Get("sound_definition_id"))
			},
			VanillaData: vanilla.SoundDefinition,
		},
	}
	Attachable.MolangLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:attachable/description/scripts/animate/*/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/initialize/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/parent_setup"),
		shared.JsonValue("minecraft:attachable/description/scripts/pre_animation/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/scale"),
		shared.JsonValue("minecraft:attachable/description/render_controllers/*/*"),
	}
	Attachable.MolangSemanticLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:attachable/description/geometry/*"),
	}
}
