package store

import (
	"github.com/ink0rr/go-jsonc"
	"github.com/ink0rr/rockide/rockide/core"
)

var EntityStore = core.NewJsonStore(core.EntityGlob, []core.JsonStoreEntry{
	{
		Id:   "id",
		Path: []jsonc.Path{{"minecraft:entity/description/identifier"}},
	},
})
