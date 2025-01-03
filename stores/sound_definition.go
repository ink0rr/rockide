package stores

import "github.com/ink0rr/rockide/core"

var SoundDefinition = newJsonStore(core.SoundDefinitionGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"sound_definitions"},
	},
})
