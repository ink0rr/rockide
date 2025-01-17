package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var Animation = newJsonStore(shared.AnimationGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"animations"},
	},
})
