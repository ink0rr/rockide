package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Block = &JsonHandler{
	Pattern: shared.BlockGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.ItemId.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:block/description/identifier")},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
		},
		{
			Store: stores.BlockTag.References,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:block/components/*"),
				shared.JsonKey("minecraft:block/permutations/*/components/*"),
			},
			Matcher: func(ctx *JsonContext) bool {
				return strings.HasPrefix(ctx.NodeValue, "tag:")
			},
			Transform: func(value string) string {
				res, _ := strings.CutPrefix(value, "tag:")
				return res
			},
		},
		{
			Store: stores.Geometry.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.References.Get()
			},
		},
		{
			Store: stores.TerrainTexture.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:material_instances/*/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:material_instances/*/texture"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:loot"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:loot"),
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:block/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
		shared.JsonValue("minecraft:block/components/minecraft:geometry/bone_visibility/*"),
		shared.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*/tags"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/bone_visibility/*"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*/tags"),
		shared.JsonValue("minecraft:block/permutations/*/condition"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:block/components/minecraft:geometry"),
		shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
	},
}
