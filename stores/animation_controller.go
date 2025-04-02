package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var AnimationController = &JsonStore{
	pattern: shared.AnimationControllerGlob,
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
