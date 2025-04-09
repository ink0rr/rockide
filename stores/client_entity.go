package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var ClientEntity = &JsonStore{
	pattern: shared.ClientEntityGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/identifier/*")},
		},
		{
			Id:   "animate",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/animations/*")},
		},
		{
			Id:   "animation_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/animations/*")},
		},
		{
			Id: "animate_refs",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:client_entity/description/scripts/animate/*/*"),
			},
		},
		{
			Id:   "material",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/materials/*")},
		},
		{
			Id:   "material_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/materials/*")},
		},
		{
			Id:   "texture",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/textures/*")},
		},
		{
			Id:   "texture_path",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/textures/*")},
		},
		{
			Id:   "geometry",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/geometry/*")},
		},
		{
			Id:   "geometry_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/geometry/*")},
		},
		{
			Id: "render_controller_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/render_controllers/*"),
				shared.JsonKey("minecraft:client_entity/description/render_controllers/*/*"),
			},
		},
		{
			Id:   "spawn_egg",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/spawn_egg/texture")},
		},
		{
			Id: "particle",
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonKey("minecraft:client_entity/description/particle_emitters/*"),
			},
		},
		{
			Id: "particle_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonValue("minecraft:client_entity/description/particle_emitters/*"),
			},
		},
		{
			Id:   "sound_definition",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/sound_effects/*")},
		},
		{
			Id:   "sound_definition_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/sound_effects/*")},
		},
	},
}
