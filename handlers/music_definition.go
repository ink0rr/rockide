package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var MusicDefintion = &JsonHandler{
	Pattern: shared.MusicDefinitionGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.MusicDefinition.Source,
			Path:       []shared.JsonPath{shared.JsonKey("*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinition.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinition.Source.Get()
			},
		},
		{
			Store: stores.SoundDefinition.References,
			Path:  []shared.JsonPath{shared.JsonValue("*/event_name")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.References.Get()
			},
		},
	},
}
