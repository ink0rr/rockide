package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var ClientAnimationController = newJsonStore(shared.ClientAnimationControllerGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"animation_controllers"},
	},
	{
		Id:   "animate_refs",
		Path: []string{"animation_controllers/*/states/*/animations"},
	},
})
