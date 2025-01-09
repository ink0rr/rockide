package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var AnimationController = jsonHandler{pattern: core.AnimationControllerGlob, entries: []jsonHandlerEntry{
	{
		Path:       []string{"animation_controllers/*"},
		MatchType:  "key",
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			filtered := []core.Reference{}
			for _, ref := range stores.Entity.Get("animation_id") {
				if strings.HasPrefix(ref.Value, "controller.") {
					filtered = append(filtered, ref)
				}
			}
			return filtered
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.AnimationController.Get("id")
		},
	},
}}
