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
          value: "<target>",
          type: ParamType.entitySelector,
        },
        {
          // todo: defined entity events
          value: ["entity_events", "minecraft:on_spawned"],
          signatureValue: "<eventName>",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default event;
