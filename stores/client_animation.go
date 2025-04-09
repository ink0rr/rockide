package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var ClientAnimation = &JsonStore{
	pattern: shared.ClientAnimationGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("animations/*")},
		},
	},
}
