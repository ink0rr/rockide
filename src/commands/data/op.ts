import { CommandInfo, ParamType } from "../types";

const op: CommandInfo = {
  command: "op",
  documentation: "Grants operator status to a player.",
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
export default op;
