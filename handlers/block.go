package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
)

var Block = &JsonHandler{Pattern: shared.BlockGlob}

func init() {
	Block.Entries = []JsonEntry{
		{
			Id:         "id",
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:block/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(ClientBlock.Get("id"), Feature.Get("block_id"))
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return Block.Get("id")
			},
		},
		{
			Id: "tag",
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
			Id: "geometry_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return Geometry.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Attachable.Get("geometry_id"), Block.Get("geometry_id"), ClientEntity.Get("geometry_id"))
			},
		},
		{
			Id: "texture_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:material_instances/*/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:material_instances/*/texture"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return TerrainTexture.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(Block.Get("texture_id"), ClientBlock.Get("texture_id"))
			},
		},
	}
	Block.MolangLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:block/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
		shared.JsonValue("minecraft:block/components/minecraft:geometry/bone_visibility/*"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/bone_visibility/*"),
		shared.JsonValue("minecraft:block/permutations/*/condition"),
	}
	Block.MolangSemanticLocations = []shared.JsonPath{
		shared.JsonValue("minecraft:block/components/minecraft:geometry"),
		shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
	}
}
