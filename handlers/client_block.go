package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientBlock = newJsonHandler(shared.ClientBlockGlob, []jsonHandlerEntry{
	{
		Path:       []shared.JsonPath{shared.JsonKey("*")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.Block.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientBlock.Get("id")
		},
	},
	{
		Path: []shared.JsonPath{
			shared.JsonValue("*/textures"),
			shared.JsonValue("*/textures/*"),
		},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.TerrainTexture.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.Block.Get("texture_id"), stores.ClientBlock.Get("texture_id"))
		},
	},
})
