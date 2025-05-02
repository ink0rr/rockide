package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/vanilla"
)

var SoundDefinition = &jsonHandler{
	pattern: shared.SoundDefinitionGlob,
	entries: []jsonHandlerEntry{
		{
			Path:       []shared.JsonPath{shared.JsonKey("sound_definitions/*")},
			Actions:    completions | definitions | rename,
			FilterDiff: true,
			Source: func(params *jsonParams) []core.Reference {
				return stores.ClientSound.Get("sound_id")
			},
			References: func(params *jsonParams) []core.Reference {
				return stores.SoundDefinition.Get("id")
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("sound_definitions/*/sounds/*"),
				shared.JsonValue("sound_definitions/*/sounds/*/name"),
			},
			Actions: completions | definitions,
			Source: func(params *jsonParams) []core.Reference {
				return stores.Sound.GetPaths()
			},
			References: func(params *jsonParams) []core.Reference {
				return nil
			},
			VanillaData: vanilla.SoundPaths,
		},
	},
}
