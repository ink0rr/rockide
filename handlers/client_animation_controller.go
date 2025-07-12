package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
)

var ClientAnimationController = &JsonHandler{Pattern: shared.ClientAnimationControllerGlob}

func init() {
	ClientAnimationController.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonKey("animation_controllers/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range slices.Concat(Attachable.Get("animation_id"), ClientEntity.Get("animation_id")) {
					if strings.HasPrefix(ref.Value, "controller.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ClientAnimationController.Get("id")
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
				return animationControllerSources(id, Attachable, ClientEntity)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return slices.Concat(
					ClientAnimationController.GetFrom(ctx.URI, "animate_refs", id),
					animationControllerReferences(id, Attachable, ClientEntity),
				)
			},
		},
	}
	ClientAnimationController.MolangLocations = []shared.JsonPath{
		shared.JsonValue("animation_controllers/*/states/*/animations/*/*"),
		shared.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_entry/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_exit/*"),
		shared.JsonValue("animation_controllers/*/states/*/parameters/*"),
		shared.JsonValue("animation_controllers/*/states/*/particle_effects/*/pre_effect_script"),
		shared.JsonValue("animation_controllers/*/states/*/variables/*/input"),
	}
}
