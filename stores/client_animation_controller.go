package stores

import (
	"github.com/ink0rr/rockide/core"
)

var ClientAnimationController = newJsonStore(core.ClientAnimationControllerGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"animation_controllers"},
	},
	{
		Id:   "animate_refs",
		Path: []string{"animation_controllers/*/states/*/animations"},
	},
})
