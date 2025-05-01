package stores

import (
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
)

var LootTable = &JsonStore{
	pattern:  shared.LootTableGlob,
	savePath: true,
	entries: []jsonStoreEntry{
		{
			Id:   "item_id",
			Path: []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Transform: func(node *jsonc.Node) *string {
				nodeValue, ok := node.Value.(string)
				if !ok || node.Parent == nil {
					return nil
				}
				parent := node.Parent.Parent
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				if entryType != nil && entryType.Value == "item" {
					return &nodeValue
				}
				return nil
			},
		},
	},
}
