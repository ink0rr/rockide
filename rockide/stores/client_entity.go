package stores

import (
	"github.com/ink0rr/rockide/rockide/core"
)

var ClientEntity = newJsonStore(core.ClientEntityGlob, []jsonStoreEntry{
	{
		Id:   "id",
		path: []string{"minecraft:client_entity/description/identifier"},
	},
})
