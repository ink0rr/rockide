package stores

import (
	"github.com/ink0rr/rockide/core"
)

var ClientEntity = newJsonStore(core.ClientEntityGlob, []jsonStoreEntry{
	{
		Id:   "id",
		Path: []string{"minecraft:client_entity/description/identifier"},
	},
	{
		Id:   "animation",
		Path: []string{"minecraft:client_entity/description/animations"},
	},
	{
		Id:   "animation_id",
		Path: []string{"minecraft:client_entity/description/animations/*"},
	},
	{
		Id:   "animate",
		Path: []string{"minecraft:client_entity/description/scripts/animate"},
	},
	{
		Id:   "material",
		Path: []string{"minecraft:client_entity/description/materials"},
	},
	{
		Id:   "material_id",
		Path: []string{"minecraft:client_entity/description/materials/*"},
	},
	{
		Id:   "texture",
		Path: []string{"minecraft:client_entity/description/textures"},
	},
	{
		Id:   "texture_path",
		Path: []string{"minecraft:client_entity/description/textures/*"},
	},
	{
		Id:   "geometry",
		Path: []string{"minecraft:client_entity/description/geometry"},
	},
	{
		Id:   "geometry_id",
		Path: []string{"minecraft:client_entity/description/geometry/*"},
	},
	{
		Id:   "render_controller_id",
		Path: []string{"minecraft:client_entity/description/render_controllers"},
	},
	{
		Id:   "spawn_egg",
		Path: []string{"minecraft:client_entity/description/spawn_egg/texture"},
	},
	{
		Id: "particle",
		Path: []string{
			"minecraft:client_entity/description/particle_effects",
			"minecraft:client_entity/description/particle_emitters",
		},
	},
	{
		Id: "particle_id",
		Path: []string{
			"minecraft:client_entity/description/particle_effects/*",
			"minecraft:client_entity/description/particle_emitters/*",
		},
	},
	{
		Id:   "sound_definition",
		Path: []string{"minecraft:client_entity/description/sound_effects"},
	},
	{
		Id:   "sound_definition_id",
		Path: []string{"minecraft:client_entity/description/sound_effects/*"},
	},
})
