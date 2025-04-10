package stores

import "github.com/ink0rr/rockide/shared"

var TerrainTexture = &JsonStore{
	pattern: shared.TerrainTextureGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("texture_data/*")},
		},
	},
}
