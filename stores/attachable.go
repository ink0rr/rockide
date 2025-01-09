package stores

import (
	"github.com/ink0rr/rockide/core"
)

var Attachable = newJsonStore(core.AttachableGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:attachable/description/identifier"},
	},
	{
		Id:   "animate",
		Path: []string{"minecraft:attachable/description/animations"},
	},
	{
		Id:   "animation_id",
		Path: []string{"minecraft:attachable/description/animations/*"},
	},
	{
		Id:   "animate_refs",
		Path: []string{"minecraft:attachable/description/scripts/animate"},
	},
	{
		Id:   "material",
		Path: []string{"minecraft:attachable/description/materials"},
	},
	{
		Id:   "material_id",
		Path: []string{"minecraft:attachable/description/materials/*"},
	},
	{
		Id:   "texture",
		Path: []string{"minecraft:attachable/description/textures"},
	},
	{
		Id:   "texture_path",
		Path: []string{"minecraft:attachable/description/textures/*"},
	},
	{
		Id:   "geometry",
		Path: []string{"minecraft:attachable/description/geometry"},
	},
	{
		Id:   "geometry_id",
		Path: []string{"minecraft:attachable/description/geometry/*"},
	},
	{
		Id:   "render_controller_id",
		Path: []string{"minecraft:attachable/description/render_controllers"},
	},
	{
		Id: "particle",
		Path: []string{
			"minecraft:attachable/description/particle_effects",
			"minecraft:attachable/description/particle_emitters",
		},
	},
	{
		Id: "particle_id",
		Path: []string{
			"minecraft:attachable/description/particle_effects/*",
			"minecraft:attachable/description/particle_emitters/*",
		},
	},
	{
		Id:   "sound_definition",
		Path: []string{"minecraft:attachable/description/sound_effects"},
	},
	{
		Id:   "sound_definition_id",
		Path: []string{"minecraft:attachable/description/sound_effects/*"},
	},
})
