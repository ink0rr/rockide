package stores

import (
	"github.com/ink0rr/rockide/core"
)

var ClientAnimation = newJsonStore(core.ClientAnimationGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"animations"},
	},
})
