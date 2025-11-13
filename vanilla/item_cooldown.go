package vanilla

import mapset "github.com/deckarep/golang-set/v2"

var ItemCooldown = mapset.NewThreadUnsafeSet(
	"chorusfruit",
	"ender_pearl",
	"wind_charge",
)
