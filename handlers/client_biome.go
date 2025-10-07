package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientBiome = &JsonHandler{
	Pattern: shared.ClientBiomeGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.BiomeId.References,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_biome/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.References.Get()
			},
		},
		{
			Store: stores.SoundDefinition.References,
			Path: sliceutil.Map([]string{
				"minecraft:ambient_sounds/addition",
				"minecraft:ambient_sounds/loop",
				"minecraft:ambient_sounds/mood",
			}, func(value string) shared.JsonPath {
				return shared.JsonValue("minecraft:client_biome/components/" + value)
			}),
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.References.Get()
			},
		},
		{
			Store: stores.Atmosphere.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:atmosphere_identifier/atmosphere_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Atmosphere.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Atmosphere.References.Get()
			},
		},
		{
			Store: stores.MusicDefinition.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:biome_music/music_definition")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinition.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinition.References.Get()
			},
		},
		{
			Store: stores.ColorGrading.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:color_grading_identifier/color_grading_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ColorGrading.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ColorGrading.References.Get()
			},
		},
		{
			Store: stores.Fog.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:fog_appearance/fog_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Fog.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Fog.References.Get()
			},
		},
		{
			Store: stores.Lighting.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:lighting_identifier/lighting_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Lighting.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Lighting.References.Get()
			},
		},
		{
			Store: stores.Water.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:water_identifier/water_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Water.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Water.References.Get()
			},
		},
	},
}
