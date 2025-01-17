package stores

import "github.com/ink0rr/rockide/shared"

var SoundDefinition = newJsonStore(shared.SoundDefinitionGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"sound_definitions"},
	},
})
