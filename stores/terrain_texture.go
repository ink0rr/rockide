package stores

import "github.com/ink0rr/rockide/core"

var TerrainTexture = newJsonStore(core.TerrainTextureGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"texture_data"},
	},
})
