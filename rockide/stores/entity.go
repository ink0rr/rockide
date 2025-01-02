package stores

import (
	"github.com/ink0rr/rockide/rockide/core"
)

var Entity = newJsonStore(core.EntityGlob, []jsonStoreEntry{
	{
		Id:   "id",
		path: []string{"minecraft:entity/description/identifier"},
	},
})
