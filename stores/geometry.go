package stores

import "github.com/ink0rr/rockide/shared"

var Geometry = &JsonStore{
	pattern: shared.GeometryGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:geometry/*/description/identifier")},
		},
	},
}
