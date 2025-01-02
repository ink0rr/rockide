package stores

import (
	"github.com/ink0rr/rockide/rockide/core"
)

var AnimationController = newJsonStore(core.AnimationControllerGlob, []jsonStoreEntry{
	{
		Id:   "id",
		path: []string{"animation_controllers"},
	},
})
