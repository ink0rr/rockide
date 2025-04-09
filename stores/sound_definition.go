package stores

import "github.com/ink0rr/rockide/shared"

var SoundDefinition = &JsonStore{
	pattern: shared.SoundDefinitionGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("sound_definitions/*")},
		},
	},
}
