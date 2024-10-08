import { CommandInfo, ParamType } from "../types";

const mobevent: CommandInfo = {
  command: "mobevent",
  documentation: "Controls what mob events are allowed to run.",
  overloads: [
    {
      params: [
        {
          value: [
            "minecraft:pillager_patrols_event",
            "minecraft:wandering_trader_event",
            "minecraft:ender_dragon_event",
            "events_enabled",
          ],
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default mobevent;
