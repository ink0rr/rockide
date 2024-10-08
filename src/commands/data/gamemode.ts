import { CommandInfo, ParamType } from "../types";

const gamemode: CommandInfo = {
  command: "gamemode",
  documentation: "Sets a player's game mode.",
  overloads: [
    {
      params: [
        {
          value: ["default", "creative", "spectator", "survival", "adventure", "d", "c", "s", "a"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
      ],
    },
    {
      params: [
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
      ],
    },
  ],
};
export default gamemode;
