import { CommandInfo, ParamType } from "../types";

const spawnpoint: CommandInfo = {
  command: "spawnpoint",
  documentation: "Sets the spawn point for a player.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default spawnpoint;
