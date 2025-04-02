package stores

import "github.com/ink0rr/rockide/shared"

var LootTable = &JsonStore{
	pattern:  shared.LootTableGlob,
	savePath: true,
	entries: []jsonStoreEntry{
		{
			Id:   "item_id",
			Path: []string{"**/entries/*/name"},
		},
	},
}
