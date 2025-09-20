package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Camera = &JsonHandler{
	Pattern: shared.CameraGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.CameraId.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:camera_preset/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.CameraId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.CameraId.Source.Get()
			},
		},
		{
			// FIXME: Prevent circular inheritance
			Store: stores.CameraId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:camera_preset/inherit_from")},
			Source: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				identifier := jsonc.FindNodeAtLocation(parent, jsonc.Path{"identifier"})
				// Prevent self-inheritance
				if identifier != nil {
					return slices.DeleteFunc(stores.CameraId.Source.Get(), func(symbol core.Symbol) bool {
						return symbol.Value == identifier.Value
					})
				}
				return stores.CameraId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.CameraId.References.Get()
			},
		},
		{
			Store: stores.AimAssistId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:camera_preset/aim_assist/preset")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.AimAssistId.References.Get()
			},
		},
	},
}
