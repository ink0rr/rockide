package stores

import (
	"github.com/ink0rr/rockide/core"
)

var ClientEntity = newJsonStore(core.ClientEntityGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:client_entity/description/identifier"},
	},
})
