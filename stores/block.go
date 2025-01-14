package stores

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
)

var Block = newJsonStore(core.BlockGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:block/description/identifier"},
	},
	{
		Id:   "tag",
		Path: []string{"minecraft:block/components", "minecraft:block/permutations/*/components"},
		Transform: func(node *jsonc.Node) *string {
			value, ok := node.Value.(string)
			if ok {
				if after, found := strings.CutPrefix(value, "tag:"); found {
					return &after
				}
			}
			return nil
		},
	},
})
