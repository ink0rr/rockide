package stores

import "github.com/ink0rr/rockide/shared"

var ItemTexture = newJsonStore(shared.ItemTextureGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"texture_data"},
	},
})
