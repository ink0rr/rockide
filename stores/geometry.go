package stores

import "github.com/ink0rr/rockide/core"

var Geometry = newJsonStore(core.GeometryGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:geometry/*/description/identifier"},
	},
})
