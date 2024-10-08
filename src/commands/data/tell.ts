import { CommandInfo, ParamType } from "../types";

const tell: CommandInfo = {
  command: "tell",
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
export default tell;
