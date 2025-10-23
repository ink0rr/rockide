package handlers

import (
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Item = &JsonHandler{
	Pattern: shared.ItemGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.ItemId.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:item/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
		},
		{
			Store: stores.ItemTexture.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:item/components/minecraft:icon"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/texture"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/textures/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTexture.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTexture.References.Get()
			},
		},
		{
			Store: stores.ItemTag.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:tags/tags/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.ItemTag.Source.Get(), stores.ItemTag.References.Get())
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/items/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:item/components/minecraft:block_placer/block"),
				shared.JsonValue("minecraft:item/components/minecraft:digger/destroy_speeds/*/block"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Store: stores.EntityId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:entity_placer/entity")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:item/components/**/condition"),
		shared.JsonValue("minecraft:item/components/minecraft:digger/destroy_speeds/*/block/tags"),
		shared.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/repair_amount"),
		shared.JsonValue("minecraft:item/components/minecraft:icon/frame"),
		shared.JsonValue("minecraft:item/events/**/sequence/*/condition"),
	},
}
