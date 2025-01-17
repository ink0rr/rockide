package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var AnimationController = newJsonHandler(shared.AnimationControllerGlob, []jsonHandlerEntry{
	{
		Matcher:    []jsonPath{matchKey("animation_controllers/*")},
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
		Matcher: []jsonPath{matchValue("animation_controllers/*/states/*/animations/*")},
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
		Matcher: []jsonPath{matchKey("animation_controllers/*/states/*/animations/*/*")},
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
})
