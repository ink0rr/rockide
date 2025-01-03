package stores

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/jsonc"
)

var Block = newJsonStore(core.BlockGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:block/description/identifier"},
	},
	{
		Id:   "tag",
		Path: []string{"minecraft:block/components", "minecraft:block/permutations/*/components"},
		Transform: func(node *jsonc.Node) transformResult {
			value, ok := node.Value.(string)
			if ok {
				if after, found := strings.CutPrefix(value, "tag:"); found {
					return transformResult{Value: after}
				}
			}
			return transformResult{Skip: true}
		},
	},
})
