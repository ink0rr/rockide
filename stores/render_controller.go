package stores

import "github.com/ink0rr/rockide/core"

var RenderController = newJsonStore(core.RenderControllerGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"render_controllers"},
	},
})
