package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var RenderController = &jsonHandler{
	pattern: shared.RenderControllerGlob,
	entries: []jsonHandlerEntry{
		{
			Path:       []shared.JsonPath{shared.JsonKey("render_controllers/*")},
			Actions:    completions | definitions | rename,
			FilterDiff: true,
			Source: func(params *jsonParams) []core.Reference {
				return slices.Concat(stores.Attachable.Get("render_controller_id"), stores.ClientEntity.Get("render_controller_id"))
			},
			References: func(params *jsonParams) []core.Reference {
				return stores.RenderController.Get("id")
			},
		},
	},
}
