package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
)

var RenderController = &JsonHandler{Pattern: shared.RenderControllerGlob}

func init() {
	RenderController.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonKey("render_controllers/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("render_controller_id"), ClientEntity.Get("render_controller_id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return RenderController.Get("id")
			},
		},
	}
	RenderController.MolangLocations = []shared.JsonPath{
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
	}
	RenderController.MolangSemanticLocations = []shared.JsonPath{
		shared.JsonValue("render_controllers/*/arrays/*/*/*"),
	}
}
