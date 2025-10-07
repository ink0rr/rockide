package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Attachable = &JsonHandler{
	Pattern: shared.AttachableGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.ItemId.References,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
		{
			Store:      stores.ClientAnimate.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/animations/*")},
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
				shared.JsonKey("minecraft:attachable/description/scripts/animate/*/*"),
				shared.JsonValue("minecraft:attachable/description/scripts/animate/*"),
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
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/animations/*")},
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
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/materials/*")},
			// TODO
		},
		{
			Store: stores.Material.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/materials/*")},
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
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/textures/*")},
			// TODO
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/textures/*")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/geometry/*")},
			// TODO
		},
		{
			Store: stores.Geometry.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/geometry/*")},
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
				shared.JsonKey("minecraft:attachable/description/render_controllers/*/*"),
				shared.JsonValue("minecraft:attachable/description/render_controllers/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/particle_effects/*"),
				shared.JsonKey("minecraft:attachable/description/particle_emitters/*"),
			},
			// TODO
		},
		{
			Store: stores.ParticleId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:attachable/description/particle_effects/*"),
				shared.JsonValue("minecraft:attachable/description/particle_emitters/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/sound_effects/*")},
			// TODO
		},
		{
			Store: stores.SoundDefinition.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/sound_effects/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.References.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:attachable/description/scripts/animate/*/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/initialize/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/parent_setup"),
		shared.JsonValue("minecraft:attachable/description/scripts/pre_animation/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/scale"),
		shared.JsonValue("minecraft:attachable/description/render_controllers/*/*"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:attachable/description/geometry/*"),
	},
}
