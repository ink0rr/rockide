package vanilla

import mapset "github.com/deckarep/golang-set/v2"

var WorldgenTemplatePool = mapset.NewThreadUnsafeSet(
	"minecraft:trail_ruins/buildings",
	"minecraft:trail_ruins/buildings/grouped",
	"minecraft:trail_ruins/decor",
	"minecraft:trail_ruins/roads",
	"minecraft:trail_ruins/tower",
	"minecraft:trail_ruins/tower/additions",
	"minecraft:trail_ruins/tower/tower_top",
)
