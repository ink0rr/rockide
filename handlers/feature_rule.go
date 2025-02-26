package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var FeatureRule = newJsonHandler(shared.FeatureRuleGlob, []jsonHandlerEntry{
	{
		Path:    []jsonPath{matchValue("minecraft:feature_rules/description/places_feature")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Feature.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Feature.Get("feature_id"), stores.FeatureRule.Get("feature_id"))
		},
	},
})
