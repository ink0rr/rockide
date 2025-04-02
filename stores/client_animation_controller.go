package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var ClientAnimationController = &JsonStore{
	pattern: shared.ClientAnimationControllerGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []string{"animation_controllers"},
		},
		{
			Id:   "animate_refs",
			Path: []string{"animation_controllers/*/states/*/animations"},
		},
	},
}
