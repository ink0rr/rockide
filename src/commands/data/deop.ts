import { CommandInfo, ParamType } from "../types";

const deop: CommandInfo = {
  command: "deop",
  documentation: "Revokes operator status from a player.",
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
export default deop;
