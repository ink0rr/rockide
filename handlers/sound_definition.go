package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var SoundDefinition = &JsonHandler{Pattern: shared.SoundDefinitionGlob}

func init() {
	SoundDefinition.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonKey("sound_definitions/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return ClientSound.Get("sound_id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return SoundDefinition.Get("id")
			},
		},
		{
			Id: "sound_path",
			Path: []shared.JsonPath{
				shared.JsonValue("sound_definitions/*/sounds/*"),
				shared.JsonValue("sound_definitions/*/sounds/*/name"),
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return Sound.GetPaths()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			VanillaData: vanilla.SoundPaths,
		},
	}
}
