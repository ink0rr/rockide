package stores

import "github.com/ink0rr/rockide/shared"

var FeatureRule = newJsonStore(shared.FeatureRuleGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:feature_rules/description/identifier"},
	},
	{
		Id:   "feature_id",
		Path: []string{"minecraft:feature_rules/description/places_feature"},
	},
})
