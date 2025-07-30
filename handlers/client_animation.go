package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientAnimation = &JsonHandler{
	Pattern: shared.ClientAnimationGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.ClientAnimation.Source,
			Path:       []shared.JsonPath{shared.JsonKey("animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range stores.ClientAnimation.References.Get() {
					if strings.HasPrefix(ref.Value, "animation.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimation.Source.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("animations/*/anim_time_update"),
		shared.JsonValue("animations/*/blend_weight"),
		shared.JsonValue("animations/*/bones/*/rotation/*"),
		shared.JsonValue("animations/*/bones/*/rotation/*/*"),
		shared.JsonValue("animations/*/bones/*/scale"),
		shared.JsonValue("animations/*/bones/*/scale/*"),
		shared.JsonValue("animations/*/bones/*/scale/*/*"),
		shared.JsonValue("animations/*/bones/*/position/*"),
		shared.JsonValue("animations/*/bones/*/position/*/*"),
		shared.JsonValue("animations/*/loop_delay"),
		shared.JsonValue("animations/*/particle_effects/*/pre_effect_script"),
		shared.JsonValue("animations/*/start_delay"),
		shared.JsonValue("animations/*/timeline/*"),
	},
}
