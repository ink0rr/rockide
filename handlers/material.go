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
			Store: stores.EntityMaterial.Source,
			Path:  []shared.JsonPath{shared.JsonKey("materials/*")},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityMaterial.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityMaterial.Source.Get()
			},
		},
		// TODO: Add base entity materials for inheritance.
	},
}

var ParticleMaterial = &JsonHandler{
	Pattern: shared.ParticleMaterialGlob,
	Entries: []JsonEntry{
		{
			Store: stores.ParticleMaterial.Source,
			Path:  []shared.JsonPath{shared.JsonKey("materials/*")},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleMaterial.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleMaterial.Source.Get()
			},
		},
		// TODO: Add base particle materials for inheritance.
	},
}
