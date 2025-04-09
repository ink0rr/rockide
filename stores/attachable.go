package stores

import (
	"github.com/ink0rr/rockide/shared"
)

var Attachable = &JsonStore{
	pattern: shared.AttachableGlob,
	entries: []jsonStoreEntry{
		{
			Id:   "id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/identifier")},
		},
		{
			Id:   "animate",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/animations/*")},
		},
		{
			Id:   "animation_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/animations/*")},
		},
		{
			Id: "animate_refs",
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/scripts/animate/*/*"),
				shared.JsonValue("minecraft:attachable/description/scripts/animate/*"),
			},
		},
		{
			Id:   "material",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/materials/*")},
		},
		{
			Id:   "material_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/materials/*")},
		},
		{
			Id:   "texture",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/textures/*")},
		},
		{
			Id:   "texture_path",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/textures/*")},
		},
		{
			Id:   "geometry",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/geometry/*")},
		},
		{
			Id:   "geometry_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/geometry/*")},
		},
		{
			Id: "render_controller_id",
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/render_controllers/*/*"),
				shared.JsonValue("minecraft:attachable/description/render_controllers/*"),
			},
		},
		{
			Id: "particle",
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/particle_effects/*"),
				shared.JsonKey("minecraft:attachable/description/particle_emitters/*"),
			},
		},
		{
			Id: "particle_id",
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:attachable/description/particle_effects/*"),
				shared.JsonValue("minecraft:attachable/description/particle_emitters/*"),
			},
		},
		{
			Id:   "sound_definition",
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/sound_effects/*")},
		},
		{
			Id:   "sound_definition_id",
			Path: []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/sound_effects/*")},
		},
	},
}
