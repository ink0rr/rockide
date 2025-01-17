package stores

import "github.com/ink0rr/rockide/shared"

var TerrainTexture = newJsonStore(shared.TerrainTextureGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"texture_data"},
	},
})
