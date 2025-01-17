package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var ClientAnimation = newJsonStore(shared.ClientAnimationGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"animations"},
	},
})
