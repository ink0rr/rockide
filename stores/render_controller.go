package stores

import "github.com/ink0rr/rockide/shared"

var RenderController = &JsonStore{
	pattern: shared.RenderControllerGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []string{"render_controllers"},
		},
	},
}
