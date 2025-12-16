package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var ClientSound = &JsonHandler{
	Pattern: shared.ClientSoundGlob,
	Entries: []JsonEntry{
		{
			Store: stores.SoundDefinition.References,
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
				return stores.SoundDefinition.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinition.References.Get()
			},
		},
		{
			Store: stores.EntityId.References,
			Path: []shared.JsonPath{
				shared.JsonKey("entity_sounds/entities/*"),
				shared.JsonKey("interactive_sounds/entity_sounds/entities/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("entity_sounds/entities/*/variants/key"),
	},
}
