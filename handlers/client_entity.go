package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientEntity = &JsonHandler{
	Pattern: shared.ClientEntityGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.EntityId.References,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
		},
		{
			Store:      stores.ClientAnimate.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimate.References.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimate.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.ClientAnimate.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:client_entity/description/scripts/animate/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimate.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimate.References.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.ClientAnimation.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/animations/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return ctx.NodeValue
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimation.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimation.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/materials/*")},
			// TODO
		},
		{
			Store: stores.Material.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/materials/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return "entity"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Material.Source.Get("entity")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Material.References.Get("entity")
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/textures/*")},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/textures/*")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/geometry/*")},
			// TODO
		},
		{
			Store: stores.Geometry.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/geometry/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.References.Get()
			},
		},
		{
			Store: stores.RenderControllerId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/render_controllers/*"),
				shared.JsonKey("minecraft:client_entity/description/render_controllers/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.References.Get()
			},
		},
		{
			Store: stores.ItemTexture.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/spawn_egg/texture")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTexture.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTexture.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonKey("minecraft:client_entity/description/particle_emitters/*"),
			},
			// TODO
		},
		{
			Store: stores.ParticleId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonValue("minecraft:client_entity/description/particle_emitters/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/sound_effects/*")},
			// TODO
		},
		{
			Store: stores.SoundDefinition.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/sound_effects/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.References.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:client_entity/description/scripts/animate/*/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/initialize/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/parent_setup"),
		shared.JsonValue("minecraft:client_entity/description/scripts/pre_animation/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/scale"),
		shared.JsonValue("minecraft:client_entity/description/render_controllers/*/*"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:client_entity/description/geometry/*"),
	},
}
