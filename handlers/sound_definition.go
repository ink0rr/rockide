package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var SoundDefinition = &JsonHandler{
	Pattern: shared.SoundDefinitionGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.SoundDefinition.Source,
			Path:       []shared.JsonPath{shared.JsonKey("sound_definitions/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.Source.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("sound_definitions/*/sounds/*"),
				shared.JsonValue("sound_definitions/*/sounds/*/name"),
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundPath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
