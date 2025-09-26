package handlers

import (
	"path/filepath"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var TextureSet = &JsonHandler{
	Pattern:   shared.TextureSetGlob,
	PathStore: stores.TexturePath,
	Entries: []JsonEntry{
		{
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:texture_set/color"),
				shared.JsonValue("minecraft:texture_set/normal"),
				shared.JsonValue("minecraft:texture_set/heightmap"),
				shared.JsonValue("minecraft:texture_set/metalness_emissive_roughness"),
				shared.JsonValue("minecraft:texture_set/metalness_emissive_roughness_subsurface"),
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				textures := stores.TexturePath.Get()
				var result []core.Symbol
				for _, texture := range textures {
					if ctx.URI.Dir().Encloses(texture.URI) && !strings.HasSuffix(texture.URI.Path(), ".json") {
						relativePath, err := filepath.Rel(ctx.URI.DirPath(), texture.URI.Path())
						if err != nil {
							continue
						}
						relativePath = filepath.ToSlash(relativePath)
						ext := filepath.Ext(relativePath)
						texture.Value = strings.TrimSuffix(relativePath, ext)
						result = append(result, texture)
					}
				}
				return result
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
