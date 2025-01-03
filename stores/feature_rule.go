package stores

import "github.com/ink0rr/rockide/core"

var FeatureRule = newJsonStore(core.FeatureRuleGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:feature_rules/description/identifier"},
	},
	{
		Id:   "feature_id",
		Path: []string{"minecraft:feature_rules/description/places_feature"},
	},
})
