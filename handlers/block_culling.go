package handlers

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var BlockCulling = &JsonHandler{
	Pattern: shared.BlockCullingGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.BlockCulling.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:block_culling_rules/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockCulling.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockCulling.Source.Get()
			},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:block_culling_rules/rules/*/geometry_part/bone")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				geos := _bcGeoIdentifiers(ctx)
				if geos == nil {
					return nil
				}
				return sliceutil.FlatMap(geos, func(identifier string) []core.Symbol {
					return stores.GeometryBone.Source.Get(identifier)
				})
			},
			References: func(ctx *JsonContext) []core.Symbol {
				geos := _bcGeoIdentifiers(ctx)
				if geos == nil {
					return nil
				}
				return sliceutil.FlatMap(geos, func(identifier string) []core.Symbol {
					return stores.GeometryBone.References.Get(identifier)
				})
			},
		},
	},
}

func _bcGeoIdentifiers(ctx *JsonContext) []string {
	// Get culilng rule identifier
	root := ctx.GetRootNode()
	identifierNode := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:block_culling_rules", "description", "identifier"})
	if identifierNode == nil {
		return nil
	}
	identifier := identifierNode.Value
	// Find geometries using the culling rules
	var identifiers []string
	set := mapset.NewThreadUnsafeSet[string]()
	for _, ref := range stores.BlockCulling.References.Get() {
		document, err := textdocument.GetOrReadFile(ref.URI)
		if err != nil {
			continue
		}
		root, _ := jsonc.ParseTree(document.GetText(), nil)
		var nodes []*jsonc.Node
		// From components
		geometryNode := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:block", "components", "minecraft:geometry"})
		if geometryNode != nil {
			nodes = append(nodes, geometryNode)
		}
		// From permutations
		permutations := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:block", "permutations"})
		if permutations != nil && permutations.Type == jsonc.NodeTypeArray {
			for _, child := range permutations.Children {
				geometryNode := jsonc.FindNodeAtLocation(child, jsonc.Path{"components", "minecraft:geometry"})
				if geometryNode != nil {
					nodes = append(nodes, geometryNode)
				}
			}
		}
		// Check each geometry node
		for _, geometryNode := range nodes {
			identifierNode := jsonc.FindNodeAtLocation(geometryNode, jsonc.Path{"identifier"})
			if identifierNode == nil {
				continue
			}
			cullingRulesNode := jsonc.FindNodeAtLocation(geometryNode, jsonc.Path{"culling"})
			if cullingRulesNode == nil {
				continue
			}
			cullingIdentifier, ok := cullingRulesNode.Value.(string)
			if !ok || cullingIdentifier != identifier {
				continue
			}
			geometryIdentifier, ok := identifierNode.Value.(string)
			if !ok {
				continue
			}
			if !set.Contains(geometryIdentifier) {
				set.Add(geometryIdentifier)
				identifiers = append(identifiers, geometryIdentifier)
			}
		}
	}
	return identifiers
}
