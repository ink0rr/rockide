package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var ClientAnimation = newJsonHandler(core.ClientAnimationGlob, []jsonHandlerEntry{
	{
		Matcher:    []jsonPath{matchKey("animations/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			filtered := []core.Reference{}
			for _, ref := range slices.Concat(stores.Attachable.Get("animation_id"), stores.ClientEntity.Get("animation_id")) {
				if strings.HasPrefix(ref.Value, "animation.") {
					filtered = append(filtered, ref)
				}
			}
			return filtered
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientAnimation.Get("id")
		},
	},
})
