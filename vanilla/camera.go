package vanilla

import mapset "github.com/deckarep/golang-set/v2"

var CameraId = mapset.NewThreadUnsafeSet(
	"minecraft:first_person",
	"minecraft:third_person",
	"minecraft:third_person_front",
	"minecraft:free",
	"minecraft:follow_orbit",
	"minecraft:fixed_boom",
	"minecraft:control_scheme_camera",
)
