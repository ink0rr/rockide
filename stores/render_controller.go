package stores

import "github.com/ink0rr/rockide/shared"

var RenderController = newJsonStore(shared.RenderControllerGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"render_controllers"},
	},
})
