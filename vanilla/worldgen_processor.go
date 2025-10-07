package vanilla

import mapset "github.com/deckarep/golang-set/v2"

var WorldgenProcessor = mapset.NewThreadUnsafeSet(
	"minecraft:trail_ruins_houses_archaeology",
	"minecraft:trail_ruins_roads_archaeology",
	"minecraft:trail_ruins_tower_top_archaeology",
)
