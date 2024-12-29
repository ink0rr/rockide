package store

import (
	"github.com/ink0rr/go-jsonc"
	"github.com/ink0rr/rockide/rockide/core"
)

var ClientEntityStore = core.NewJsonStore(core.ClientEntityGlob, []core.JsonStoreEntry{
	{
		Id:   "id",
		Path: []jsonc.Path{{"minecraft:client_entity/description/identifier"}},
	},
})
