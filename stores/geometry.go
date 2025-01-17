package stores

import "github.com/ink0rr/rockide/shared"

var Geometry = newJsonStore(shared.GeometryGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:geometry/*/description/identifier"},
	},
})
