import { CommandInfo, ParamType } from "../types";

const clearspawnpoint: CommandInfo = {
  command: "clearspawnpoint",
  documentation: "Removes the spawn point for a player.",
  overloads: [
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
      ],
    },
  ],
};
export default clearspawnpoint;
