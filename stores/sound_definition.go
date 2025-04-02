package stores

import "github.com/ink0rr/rockide/shared"

var SoundDefinition = &JsonStore{
	pattern: shared.SoundDefinitionGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []string{"sound_definitions"},
		},
	},
}
