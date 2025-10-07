package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var WorldgenStructureSet = &JsonHandler{
	Pattern: shared.WorldgenStructureSetGlob,
	Entries: []JsonEntry{
		{
			Store: stores.WorldgenJigsaw.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:structure_set/structures/*/structure")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsaw.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsaw.References.Get()
			},
		},
	},
}
