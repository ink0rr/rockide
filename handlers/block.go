package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
	"golang.org/x/mod/semver"
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
			Store:      stores.BlockState.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:block/description/states/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockState.References.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockState.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.BlockTag.Source,
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
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BlockTag.Source.Get(), stores.BlockTag.References.Get())
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.Geometry.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
				shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry"),
				shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/identifier"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry/identifier"),
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
				shared.JsonValue("minecraft:block/components/minecraft:item_visual/material_instances/*/texture"),
				shared.JsonValue("minecraft:block/components/minecraft:destruction_particles/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:material_instances/*/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/material_instances/*/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:destruction_particles/texture"),
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
		{
			Store: stores.BlockCulling.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:geometry/culling"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/culling"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockCulling.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockCulling.References.Get()
			},
		},
		{
			Store: stores.RecipeTag.Source,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:crafting_table/crafting_tags/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:crafting_table/crafting_tags/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.RecipeTag.Source.Get(), stores.RecipeTag.References.Get())
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/block_filter/*"),
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/block_filter/*/name"),
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*"),
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*/name"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/block_filter/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/block_filter/*/name"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*/name"),
			},
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
			Store: stores.BlockCustomComponent.Source,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:custom_components/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:custom_components/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BlockCustomComponent.Source.Get(), stores.BlockCustomComponent.References.Get())
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.BlockCustomComponent.Source,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:block/components/*"),
				shared.JsonKey("minecraft:block/permutations/*/components/*"),
			},
			Matcher: func(ctx *JsonContext) bool {
				root := ctx.GetRootNode()
				formatNode := jsonc.FindNodeAtLocation(root, jsonc.Path{"format_version"})
				if formatNode != nil {
					if version, ok := formatNode.Value.(string); ok {
						match := !strings.HasPrefix(ctx.NodeValue, "minecraft:") && !strings.HasPrefix(ctx.NodeValue, "tag:")
						return semver.Compare("v"+version, "v1.21.80") > 0 && match
					}
				}
				return false
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BlockCustomComponent.Source.Get(), stores.BlockCustomComponent.References.Get())
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
		shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry"),
		shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/identifier"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry/identifier"),
	},
}
