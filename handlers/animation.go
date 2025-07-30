package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Animation = &JsonHandler{
	Pattern: shared.AnimationGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.Animation.Source,
			Path:       []shared.JsonPath{shared.JsonKey("animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range stores.Animation.References.Get() {
					if strings.HasPrefix(ref.Value, "animation.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Animation.Source.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("animations/*/anim_time_update"),
		shared.JsonValue("animations/*/timeline/*/*"),
	},
}
