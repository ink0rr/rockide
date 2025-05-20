package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

var ClientSound = &JsonHandler{Pattern: shared.ClientSoundGlob}

func init() {
	ClientSound.Entries = []JsonEntry{
		{
			Id: "sound_id",
			Path: []shared.JsonPath{
				shared.JsonValue("block_sounds/*/events/*"),
				shared.JsonValue("block_sounds/*/events/*/sound"),
				shared.JsonValue("entity_sounds/entities/*/events/*"),
				shared.JsonValue("entity_sounds/entities/*/events/*/sound"),
				shared.JsonValue("entity_sounds/entities/*/variants/map/*/events/*"),
				shared.JsonValue("entity_sounds/entities/*/variants/map/*/events/*/sound"),
				shared.JsonValue("individual_event_sounds/events/*"),
				shared.JsonValue("individual_event_sounds/events/*/sound"),
				shared.JsonValue("individual_named_sounds/sounds/*"),
				shared.JsonValue("individual_named_sounds/sounds/*/sound"),
				shared.JsonValue("interactive_sounds/*/*/events/*"),
				shared.JsonValue("interactive_sounds/*/*/events/*/sound"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return SoundDefinition.Get("id")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return ClientSound.Get("sound_id")
			},
			VanillaData: vanilla.SoundDefinition,
		},
	}
	ClientSound.MolangLocations = []shared.JsonPath{
		shared.JsonValue("entity_sounds/entities/*/variants/key"),
	}
}
