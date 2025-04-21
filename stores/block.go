package stores

import (
	"strings"

	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
)

var Block = &JsonStore{
	pattern: shared.BlockGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:block/description/identifier")},
		},
		{
			Id: "tag",
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:block/components/*"),
				shared.JsonKey("minecraft:block/permutations/*/components/*"),
			},
			Transform: func(node *jsonc.Node) *string {
				nodeValue, ok := node.Value.(string)
				if ok {
					if after, found := strings.CutPrefix(nodeValue, "tag:"); found {
						return &after
					}
				}
				return nil
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
		},
	},
}
