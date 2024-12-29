package store

import (
	"github.com/ink0rr/go-jsonc"
	"github.com/ink0rr/rockide/rockide/core"
)

var AnimationControllerStore = core.NewJsonStore(core.AnimationControllerGlob, []core.JsonStoreEntry{
	{
		Id:   "id",
		Path: []jsonc.Path{{"animation_controllers"}},
	},
})
