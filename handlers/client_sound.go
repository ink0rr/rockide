package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/vanilla"
)

var ClientSound = &jsonHandler{
	pattern: shared.ClientSoundGlob,
	entries: []jsonHandlerEntry{
		{
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
			Actions: completions | definitions | rename,
			Source: func(params *jsonParams) []core.Reference {
				return stores.SoundDefinition.Get("id")
			},
			References: func(params *jsonParams) []core.Reference {
				return stores.ClientSound.Get("sound_id")
			},
			VanillaData: vanilla.SoundDefinition,
		},
	},
}
