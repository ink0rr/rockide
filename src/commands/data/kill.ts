import { CommandInfo, ParamType } from "../types";

const kill: CommandInfo = {
  command: "kill",
  documentation: "Kills entities (players, mobs, etc.).",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
      ],
    },
  ],
};
export default kill;
