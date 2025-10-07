package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var WorldgenTemplatePool = &JsonHandler{
	Pattern: shared.WorldgenTemplatePoolGlob,
	Entries: []JsonEntry{
		{
			Store: stores.WorldgenTemplatePool.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:template_pool/description/identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePool.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePool.Source.Get()
			},
		},
		{
			Store: stores.WorldgenProcessor.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:template_pool/elements/*/element/processors")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenProcessor.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenProcessor.References.Get()
			},
		},
		{
			Store: stores.WorldgenTemplatePool.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:template_pool/fallback")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePool.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePool.References.Get()
			},
		},
		// TODO: Structure references
	},
}
