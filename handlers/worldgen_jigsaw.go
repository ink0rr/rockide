package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var WorldgenJigsaw = &JsonHandler{
	Pattern: shared.WorldgenJigsawGlob,
	Entries: []JsonEntry{
		{
			Store: stores.WorldgenJigsaw.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:jigsaw/description/identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsaw.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsaw.Source.Get()
			},
		},
		{
			Store: stores.BiomeTag.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:jigsaw/biome_filters/**/value")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.References.Get()
			},
		},
		{
			Store: stores.WorldgenTemplatePool.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:jigsaw/start_pool"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/alias"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/target"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/targets/*/data"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/alias"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/target"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePool.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePool.References.Get()
			},
		},
	},
}
