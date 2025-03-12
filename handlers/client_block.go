package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientBlock = newJsonHandler(shared.ClientBlockGlob, []jsonHandlerEntry{
	{
		Path:       []jsonPath{matchKey("*")},
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
		Path:    []jsonPath{matchValue("*/textures"), matchValue("*/textures/*")},
		Actions: completions | definitions | rename,
		Source: func(params *jsonParams) []core.Reference {
			return stores.TerrainTexture.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.ClientBlock.Get("texture")
		},
	},
})
