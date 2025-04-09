package stores

import (
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
)

var ClientBlock = &JsonStore{
	pattern: shared.ClientBlockGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("*")},
			Transform: func(node *jsonc.Node) *string {
				nodeValue, ok := node.Value.(string)
				if !ok || nodeValue == "format_version" {
					return nil
				}
				return &nodeValue
			},
		},
		{
			Id: "texture",
			Path: []shared.JsonPath{
				shared.JsonValue("*/textures"),
				shared.JsonValue("*/textures/*"),
			},
		},
	},
}
