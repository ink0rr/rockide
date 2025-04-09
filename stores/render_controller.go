package stores

import "github.com/ink0rr/rockide/shared"

var RenderController = &JsonStore{
	pattern: shared.RenderControllerGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("render_controllers/*")},
		},
	},
}
