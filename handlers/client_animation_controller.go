package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientAnimationController = newJsonHandler(shared.ClientAnimationControllerGlob, []jsonHandlerEntry{
	{
		Path:       []jsonPath{matchKey("animation_controllers/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			filtered := []core.Reference{}
			for _, ref := range slices.Concat(stores.Attachable.Get("animation_id"), stores.ClientEntity.Get("animation_id")) {
				if strings.HasPrefix(ref.Value, "controller.") {
					filtered = append(filtered, ref)
				}
			}
			return filtered
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientAnimationController.Get("id")
		},
	},
	{
		Path:    []jsonPath{matchValue("animation_controllers/*/states/*/animations/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerSources(id, stores.Attachable, stores.ClientEntity)
		},
		References: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerReferences(id, stores.ClientAnimationController, stores.Attachable, stores.ClientEntity)
		},
	},
	{
		Path:    []jsonPath{matchKey("animation_controllers/*/states/*/animations/*/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerSources(id, stores.Attachable, stores.ClientEntity)
		},
		References: func(params *jsonParams) []core.Reference {
			id, ok := params.Location.Path[1].(string)
			if !ok {
				return nil
			}
			return animationControllerReferences(id, stores.ClientAnimationController, stores.Attachable, stores.ClientEntity)
		},
	},
})
