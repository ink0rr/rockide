package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
)

var AnimationController = &JsonHandler{Pattern: shared.AnimationControllerGlob}

func init() {
	AnimationController.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonKey("animation_controllers/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range Entity.Get("animation_id") {
					if strings.HasPrefix(ref.Value, "controller.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return AnimationController.Get("id")
			},
		},
		{
			Id: "animate_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("animation_controllers/*/states/*/animations/*"),
				shared.JsonKey("animation_controllers/*/states/*/animations/*/*"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				if id, ok := ctx.GetPath()[1].(string); ok {
					return id
				}
				return defaultScope
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return animationControllerSources(id, Entity)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return slices.Concat(
					AnimationController.GetFrom(ctx.URI, "animate_refs", id),
					animationControllerReferences(id, Entity),
				)
			},
		},
	}
	AnimationController.MolangLocations = []shared.JsonPath{
		shared.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_entry/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_exit/*"),
	}
}
