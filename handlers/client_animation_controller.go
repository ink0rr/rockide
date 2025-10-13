package handlers

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientAnimationController = &JsonHandler{
	Pattern: shared.ClientAnimationControllerGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.ClientAnimation.Source,
			Path:       []shared.JsonPath{shared.JsonKey("animation_controllers/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range stores.ClientAnimation.References.Get() {
					if strings.HasPrefix(ref.Value, "controller.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimation.Source.Get()
			},
		},
		{
			Store: stores.ClientAnimate.References,
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
				res := []core.Symbol{}
				set := mapset.NewThreadUnsafeSet[protocol.DocumentURI]()
				for _, symbol := range stores.ClientAnimation.References.Get(id) {
					if !set.ContainsOne(symbol.URI) {
						set.Add(symbol.URI)
						res = append(res, stores.ClientAnimate.Source.GetFrom(symbol.URI)...)
					}
				}
				return res
			},
			References: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				res := []core.Symbol{}
				set := mapset.NewThreadUnsafeSet[protocol.DocumentURI]()
				for _, symbol := range stores.ClientAnimation.References.Get(id) {
					if !set.ContainsOne(symbol.URI) {
						set.Add(symbol.URI)
						res = append(res, stores.ClientAnimate.References.GetFrom(symbol.URI)...)
					}
				}
				return res
			},
		},
		{
			Store:      stores.ClientControllerState.Source,
			Path:       []shared.JsonPath{shared.JsonKey("animation_controllers/*/states/*")},
			FilterDiff: true,
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
				return stores.ClientControllerState.References.GetFrom(ctx.URI, id)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return stores.ClientControllerState.Source.GetFrom(ctx.URI, id)
			},
		},
		{
			Store: stores.ClientControllerState.References,
			Path: []shared.JsonPath{
				shared.JsonValue("animation_controllers/*/initial_state"),
				shared.JsonKey("animation_controllers/*/states/*/transitions/*/*"),
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
				return stores.ClientControllerState.Source.GetFrom(ctx.URI, id)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return stores.ClientControllerState.References.GetFrom(ctx.URI, id)
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("animation_controllers/*/states/*/animations/*/*"),
		shared.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_entry/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_exit/*"),
		shared.JsonValue("animation_controllers/*/states/*/parameters/*"),
		shared.JsonValue("animation_controllers/*/states/*/particle_effects/*/pre_effect_script"),
		shared.JsonValue("animation_controllers/*/states/*/variables/*/input"),
	},
}
