package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
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
		{
			Store: stores.BiomeTag.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/conditions/minecraft:biome_filter/**/value")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_biome_tag"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.References.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:feature_rules/distribution/iterations"),
		shared.JsonValue("minecraft:feature_rules/distribution/scatter_chance/*"),
		shared.JsonValue("minecraft:feature_rules/distribution/x/extent/*"),
		shared.JsonValue("minecraft:feature_rules/distribution/y/extent/*"),
		shared.JsonValue("minecraft:feature_rules/distribution/z/extent/*"),
		shared.JsonValue("minecraft:feature_rules/distribution/x"),
		shared.JsonValue("minecraft:feature_rules/distribution/y"),
		shared.JsonValue("minecraft:feature_rules/distribution/z"),
	},
}
