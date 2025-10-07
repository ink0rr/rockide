package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Biome = &JsonHandler{
	Pattern: shared.BiomeGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.BiomeId.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:biome/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.Source.Get()
			},
		},
		{
			Store: stores.BiomeTag.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:biome/components/minecraft:tags/tags/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BiomeTag.Source.Get(), stores.BiomeTag.References.Get())
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.ItemId.References,
			Path: sliceutil.Map([]string{
				"minecraft:mountain_parameters/steep_material_adjustment/material",
				"minecraft:mountain_parameters/steep_material_adjustment/material/name",
				"minecraft:surface_builder/builder/foundation_material",
				"minecraft:surface_builder/builder/foundation_material/name",
				"minecraft:surface_builder/builder/mid_material",
				"minecraft:surface_builder/builder/mid_material/name",
				"minecraft:surface_builder/builder/sea_floor_material",
				"minecraft:surface_builder/builder/sea_floor_material/name",
				"minecraft:surface_builder/builder/sea_material",
				"minecraft:surface_builder/builder/sea_material/name",
				"minecraft:surface_builder/builder/top_material",
				"minecraft:surface_builder/builder/top_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/floor_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/floor_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/foundation_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/foundation_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/mid_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/mid_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_floor_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_floor_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/top_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/top_material/name",
			}, func(value string) shared.JsonPath {
				return shared.JsonValue("minecraft:biome/components/" + value)
			}),
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Store: stores.FeatureId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:biome/components/minecraft:forced_features/*/*/places_feature/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.FeatureId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.FeatureId.References.Get()
			},
		},
		// TODO: Add support for states.
	},
}
