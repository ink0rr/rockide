package stores

import "github.com/ink0rr/rockide/shared"

var FeatureRule = &JsonStore{
	pattern: shared.FeatureRuleGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/description/identifier")},
		},
		{
			Id:   "feature_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/description/places_feature")},
		},
	},
}
