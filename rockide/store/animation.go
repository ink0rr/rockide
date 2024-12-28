package store

import (
	"github.com/ink0rr/rockide/jsonc"
	"github.com/ink0rr/rockide/rockide/core"
)

var AnimationStore = core.NewJsonStore(core.AnimationGlob, []core.JsonStoreEntry{
	{
		Id:   "id",
		Path: []jsonc.Path{{"animations"}},
	},
})
