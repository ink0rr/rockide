package stores

import "github.com/ink0rr/rockide/shared"

var ItemTexture = &JsonStore{
	pattern: shared.ItemTextureGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []string{"texture_data"},
		},
	},
}
