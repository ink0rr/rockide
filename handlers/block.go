package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/stores"
)

var Block = jsonHandler{pattern: core.BlockGlob, entries: []jsonHandlerEntry{
	{
		Path:       []string{"minecraft:block/description/identifier"},
		MatchType:  "value",
		Actions:    completions | definitions | rename,
		FilterDiff: true,
		Source: func(params *jsonParams) []core.Reference {
			return stores.ClientBlock.Get("id")
		},
		References: func(params *jsonParams) []core.Reference {
			return stores.Block.Get("id")
		},
	},
}}
