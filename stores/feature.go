package stores

import (
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
)

var Feature = &JsonStore{
	pattern: shared.FeatureGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []string{"*/description/identifier"},
		},
		{
			Id: "block_id",
			Path: sliceutil.FlatMap([]string{
				"minecraft:catalyst_feature/can_place_sculk_catalyst_on/*",
				"minecraft:catalyst_feature/central_block",
				"minecraft:cave_carver_feature/fill_with",
				"minecraft:fossil_feature/ore_block",
				"minecraft:geode_feature/alternate_inner_layer",
				"minecraft:geode_feature/filler",
				"minecraft:geode_feature/inner_layer",
				"minecraft:geode_feature/inner_placements/*",
				"minecraft:geode_feature/middle_layer",
				"minecraft:geode_feature/outer_layer",
				"minecraft:growing_plant_feature/body_blocks/*/0",
				"minecraft:growing_plant_feature/head_blocks/*/0",
				"minecraft:hell_cave_carver_feature/fill_with",
				"minecraft:multiface_feature/can_place_on/*",
				"minecraft:multiface_feature/places_block",
				"minecraft:ore_feature/replace_rules/*/may_replace/*",
				"minecraft:ore_feature/replace_rules/*/places_block",
				"minecraft:partially_exposed_blob_feature/places_block",
				"minecraft:single_block_feature/may_attach_to/top/*",
				"minecraft:single_block_feature/may_attach_to/bottom/*",
				"minecraft:single_block_feature/may_attach_to/north/*",
				"minecraft:single_block_feature/may_attach_to/east/*",
				"minecraft:single_block_feature/may_attach_to/south/*",
				"minecraft:single_block_feature/may_attach_to/west/*",
				"minecraft:single_block_feature/may_attach_to/all/*",
				"minecraft:single_block_feature/may_attach_to/diagonal/*",
				"minecraft:single_block_feature/may_attach_to/sides/*",
				"minecraft:single_block_feature/may_not_attach_to/*/*",
				"minecraft:single_block_feature/may_replace/*",
				"minecraft:single_block_feature/places_block",
				"minecraft:structure_template_feature/constraints/block_intersection/block_allowlist/*",
				"minecraft:tree_feature/base_cluster/may_replace/*",
				"minecraft:tree_feature/mangrove_roots/above_root/above_root_block",
				"minecraft:tree_feature/mangrove_roots/mud_block",
				"minecraft:tree_feature/mangrove_roots/muddy_root_block",
				"minecraft:tree_feature/mangrove_roots/root_block",
				"minecraft:tree_feature/mangrove_roots/roots_may_grow_through/*",
				"minecraft:tree_feature/may_grow_on/*",
				"minecraft:tree_feature/may_grow_through/*",
				"minecraft:tree_feature/may_replace/*",
				"minecraft:tree_feature/trunk/base_block",
				"minecraft:tree_feature/trunk/trunk_block",
				"minecraft:tree_feature/**/decoration_block",
				"minecraft:tree_feature/**/hanging_block",
				"minecraft:tree_feature/**/leaf_block",
				"minecraft:tree_feature/**/leaf_blocks/*",
				"minecraft:tree_feature/**/trunk_block",
				"minecraft:underwater_cave_carver_feature/fill_with",
				"minecraft:underwater_cave_carver_feature/replace_air_with",
				"minecraft:vegetation_patch_feature/ground_block",
				"minecraft:vegetation_patch_feature/replaceable_blocks/*",
			}, func(value string) []string {
				return []string{value, value + "/name"}
			}),
		},
		{
			Id: "feature_id",
			Path: []string{
				"minecraft:aggregate_feature/features/*",
				"minecraft:catalyst_feature/central_patch_feature",
				"minecraft:catalyst_feature/patch_feature",
				"minecraft:scatter_feature/places_feature",
				"minecraft:search_feature/places_feature",
				"minecraft:sequence_feature/features/*",
				"minecraft:snap_to_surface_feature/feature_to_snap",
				"minecraft:surface_relative_threshold_feature/feature_to_place",
				"minecraft:vegetation_patch_feature/vegetation_feature",
				"minecraft:weighted_random_feature/features/*/0",
			},
		},
	},
}
