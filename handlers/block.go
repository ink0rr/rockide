package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Block = newJsonHandler(shared.BlockGlob, []jsonHandlerEntry{
	{
		Matcher:    []jsonPath{matchValue("minecraft:block/description/identifier")},
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return slices.Concat(stores.ClientBlock.Get("id"), stores.Feature.Get("block_id"))
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Block.Get("id")
		},
	},
})
