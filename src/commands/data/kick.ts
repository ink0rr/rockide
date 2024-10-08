import { CommandInfo, ParamType } from "../types";

const kick: CommandInfo = {
  command: "kick",
  documentation: "Kicks a player from the server.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "string",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default kick;
