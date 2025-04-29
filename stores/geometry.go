package stores

import (
	"strings"

	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
)

var Geometry = &JsonStore{
	pattern: shared.GeometryGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:geometry/*/description/identifier")},
		},
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonKey("*")},
			Transform: func(node *jsonc.Node) *string {
				nodeValue, ok := node.Value.(string)
				if !ok || !strings.HasPrefix(nodeValue, "geometry.") {
					return nil
				}
				before, _, found := strings.Cut(nodeValue, ":")
				if found {
					return &before
				}
				return &nodeValue
			},
		},
	},
}
