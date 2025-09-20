package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Dialogue = &JsonHandler{
	Pattern: shared.DialogueGlob,
	Entries: []JsonEntry{
		{
			Store: stores.DialogueId.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:npc_dialogue/scenes/*/scene_tag")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.DialogueId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.DialogueId.Source.Get()
			},
		},
	},
}
