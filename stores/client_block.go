package stores

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/jsonc"
)

var ClientBlock = newJsonStore(core.ClientBlockGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"*"},
		Transform: func(node *jsonc.Node) *string {
			nodeValue, ok := node.Value.(string)
			if !ok || node.Value == "format_version" {
				return nil
			}
			return &nodeValue
		},
	},
	{
		Id:        "texture",
		Path:      []string{"*/textures", "*/textures/*"},
		Transform: skipKey,
	},
})
