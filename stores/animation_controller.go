package stores

import (
	"github.com/ink0rr/rockide/core"
)

var AnimationController = newJsonStore(core.AnimationControllerGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"animation_controllers"},
	},
})
