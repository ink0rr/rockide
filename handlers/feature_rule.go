package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
)

var FeatureRule = &JsonHandler{Pattern: shared.FeatureRuleGlob}

func init() {
	FeatureRule.Entries = []JsonEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/description/identifier")},
		},
		{
			Id:   "feature_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/description/places_feature")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Feature.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Feature.Get("feature_id"), FeatureRule.Get("feature_id"))
			},
		},
	}
}
