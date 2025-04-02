package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var Animation = &JsonStore{
	pattern: shared.AnimationGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []string{"animations"},
		},
	},
}
