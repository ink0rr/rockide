package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
)

var ClientAnimation = &JsonHandler{Pattern: shared.ClientAnimationGlob}

func init() {
	ClientAnimation.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonKey("animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range slices.Concat(Attachable.Get("animation_id"), ClientEntity.Get("animation_id")) {
					if strings.HasPrefix(ref.Value, "animation.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ClientAnimation.Get("id")
			},
		},
	}
	ClientAnimation.MolangLocations = []shared.JsonPath{
		shared.JsonValue("animations/*/anim_time_update"),
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
	}
}
