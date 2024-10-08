import { CommandInfo, ParamType } from "../types";

const msg: CommandInfo = {
  command: "msg",
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
export default msg;
