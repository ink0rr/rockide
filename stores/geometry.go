package stores

import "github.com/ink0rr/rockide/shared"

var Geometry = &JsonStore{
	pattern: shared.GeometryGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []string{"minecraft:geometry/*/description/identifier"},
		},
	},
}
