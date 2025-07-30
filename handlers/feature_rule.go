package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var FeatureRule = &JsonHandler{
	Pattern: shared.FeatureRuleGlob,
	Entries: []JsonEntry{
		{
			Store: stores.FeatureId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/description/places_feature")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.FeatureId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.FeatureId.References.Get()
			},
		},
	},
}
