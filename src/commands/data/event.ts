import { CommandInfo, ParamType } from "../types";

const event: CommandInfo = {
  command: "event",
  documentation: "Triggers an event for the specified object(s)",
  overloads: [
    {
      params: [
        {
          value: ["entity"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["entity_events", "minecraft:on_spawned"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default event;
