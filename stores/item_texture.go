package stores

import "github.com/ink0rr/rockide/core"

var ItemTexture = newJsonStore(core.ItemTextureGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"texture_data"},
	},
})
