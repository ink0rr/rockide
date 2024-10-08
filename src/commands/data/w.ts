import { CommandInfo, ParamType } from "../types";

const w: CommandInfo = {
  command: "w",
  documentation: "Sends a private message to one or more players.",
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
export default w;
