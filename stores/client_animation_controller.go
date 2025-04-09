package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var ClientAnimationController = &JsonStore{
	pattern: shared.ClientAnimationControllerGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("animation_controllers/*")},
		},
		{
			Id: "animate_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("animation_controllers/*/states/*/animations/*"),
				shared.JsonKey("animation_controllers/*/states/*/animations/*/*"),
			},
		},
	},
}
