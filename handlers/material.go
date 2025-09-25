package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var EntityMaterial = &JsonHandler{
	Pattern: shared.EntityMaterialGlob,
	Entries: []JsonEntry{
		{
			Store: stores.Material.Source,
			Path:  []shared.JsonPath{shared.JsonKey("materials/*")},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				return "entity"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Material.References.Get("entity")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Material.Source.Get("entity")
			},
		},
		// TODO: Add base entity materials for inheritance.
	},
}

var ParticleMaterial = &JsonHandler{
	Pattern: shared.ParticleMaterialGlob,
	Entries: []JsonEntry{
		{
			Store: stores.Material.Source,
			Path:  []shared.JsonPath{shared.JsonKey("materials/*")},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				return "particle"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Material.References.Get("particle")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Material.Source.Get("particle")
			},
		},
		// TODO: Add base particle materials for inheritance.
	},
}
