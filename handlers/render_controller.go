package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var RenderController = &JsonHandler{
	Pattern: shared.RenderControllerGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.RenderControllerId.Source,
			Path:       []shared.JsonPath{shared.JsonKey("render_controllers/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.Source.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("render_controllers/*/uv_anim/offset/*"),
		shared.JsonValue("render_controllers/*/uv_anim/scale/*"),
		shared.JsonValue("render_controllers/*/geometry"),
		shared.JsonValue("render_controllers/*/part_visibility/*/*"),
		shared.JsonValue("render_controllers/*/materials/*/*"),
		shared.JsonValue("render_controllers/*/textures/*"),
		shared.JsonValue("render_controllers/*/color/*"),
		shared.JsonValue("render_controllers/*/overlay_color/*"),
		shared.JsonValue("render_controllers/*/is_hurt_color/*"),
		shared.JsonValue("render_controllers/*/on_fire_color/*"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("render_controllers/*/arrays/*/*/*"),
	},
}
