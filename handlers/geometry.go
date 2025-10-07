package handlers

import (
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Geometry = &JsonHandler{
	Pattern: shared.GeometryGlob,
	Entries: []JsonEntry{
		{
			Store: stores.Geometry.Source,
			Path: []shared.JsonPath{
				shared.JsonKey("*"),
				shared.JsonValue("minecraft:geometry/*/description/identifier"),
			},
			Matcher: func(ctx *JsonContext) bool {
				return strings.HasPrefix(ctx.NodeValue, "geometry.")
			},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Geometry.Source.Get()
			},
		},
		{
			Store:    stores.GeometryBone.Source,
			Path:     []shared.JsonPath{shared.JsonValue("minecraft:geometry/*/bones/*/name")},
			ScopeKey: _geometryGetIdentifier,
			Source: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetIdentifier(ctx)
				return stores.GeometryBone.References.Get(identifier)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetIdentifier(ctx)
				return stores.GeometryBone.Source.Get(identifier)
			},
		},
		{
			// FIXME: Prevent circular references
			Store:    stores.GeometryBone.References,
			Path:     []shared.JsonPath{shared.JsonValue("minecraft:geometry/*/bones/*/parent")},
			ScopeKey: _geometryGetIdentifier,
			Source: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetIdentifier(ctx)
				// Prevent self-reference
				parent := ctx.GetParentNode()
				name := jsonc.FindNodeAtLocation(parent, jsonc.Path{"name"})
				result := stores.GeometryBone.Source.Get(identifier)
				if name != nil {
					return slices.DeleteFunc(result, func(symbol core.Symbol) bool {
						return symbol.Value == name.Value
					})
				}
				return result
			},
			References: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetIdentifier(ctx)
				return stores.GeometryBone.References.Get(identifier)
			},
		},
		// Older version of geometry files
		{
			Store: stores.GeometryBone.Source,
			Path:  []shared.JsonPath{shared.JsonValue("*/bones/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				return _geometryGetOldIdentifier(ctx) != defaultScope
			},
			ScopeKey: _geometryGetOldIdentifier,
			Source: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetOldIdentifier(ctx)
				return stores.GeometryBone.References.Get(identifier)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetOldIdentifier(ctx)
				return stores.GeometryBone.Source.Get(identifier)
			},
		},
		{
			// FIXME: Prevent circular references
			Store: stores.GeometryBone.References,
			Path:  []shared.JsonPath{shared.JsonValue("*/bones/*/parent")},
			Matcher: func(ctx *JsonContext) bool {
				return _geometryGetOldIdentifier(ctx) != defaultScope
			},
			ScopeKey: _geometryGetOldIdentifier,
			Source: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetOldIdentifier(ctx)
				// Prevent self-reference
				parent := ctx.GetParentNode()
				name := jsonc.FindNodeAtLocation(parent, jsonc.Path{"name"})
				result := stores.GeometryBone.Source.Get(identifier)
				if name != nil {
					return slices.DeleteFunc(result, func(symbol core.Symbol) bool {
						return symbol.Value == name.Value
					})
				}
				return result
			},
			References: func(ctx *JsonContext) []core.Symbol {
				identifier := _geometryGetOldIdentifier(ctx)
				return stores.GeometryBone.References.Get(identifier)
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:geometry/*/bones/*/binding"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:geometry/*/description/identifier"),
	},
}

func _geometryGetIdentifier(ctx *JsonContext) string {
	path := ctx.GetPath()
	path = slices.Clone(path[:len(path)-3])
	path = append(path, "description", "identifier")
	node := jsonc.FindNodeAtLocation(ctx.GetRootNode(), jsonc.Path(path))
	if node != nil {
		if id, ok := node.Value.(string); ok {
			return id
		}
	}
	return defaultScope
}

func _geometryGetOldIdentifier(ctx *JsonContext) string {
	node := ctx.GetParentNode().Parent.Parent.Parent.Parent
	identifier, ok := node.Children[0].Value.(string)
	if ok {
		return identifier
	}
	return defaultScope
}
