package stores

import "github.com/ink0rr/rockide/shared"

var ClientSound = &JsonStore{
	pattern: shared.ClientSoundGlob,
	entries: []jsonStoreEntry{
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
		},
	},
}
