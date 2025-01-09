package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var AnimationController = newJsonHandler(core.AnimationControllerGlob, []jsonHandlerEntry{
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
	{
		Path: []string{
			"animation_controllers/*/states/*/animations/*",
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerSources(id, stores.Entity)
		},
		References: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerReferences(id, stores.AnimationController, stores.Entity)
		},
	},
	{
		Path: []string{
			"animation_controllers/*/states/*/animations/*/*",
		},
		MatchType: "key",
		Actions:   completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerSources(id, stores.Entity)
		},
		References: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerReferences(id, stores.AnimationController, stores.Entity)
		},
	},
})
