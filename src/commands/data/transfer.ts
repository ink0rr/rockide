import { CommandInfo, ParamType } from "../types";

const transfer: CommandInfo = {
  command: "transfer",
  documentation: "Transfers a player to another server.",
  overloads: [
    {
      params: [
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default transfer;
