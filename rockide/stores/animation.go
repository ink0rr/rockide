package stores

import (
	"github.com/ink0rr/rockide/rockide/core"
)

var Animation = newJsonStore(core.AnimationGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"animations"},
	},
})
