package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var Animation = newJsonHandler(core.AnimationGlob, []jsonHandlerEntry{
	{
		Matcher:    []jsonPath{matchKey("animations/*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			filtered := []core.Reference{}
			for _, ref := range stores.Entity.Get("animation_id") {
				if strings.HasPrefix(ref.Value, "animation.") {
					filtered = append(filtered, ref)
				}
			}
			return filtered
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Animation.Get("id")
		},
	},
})
