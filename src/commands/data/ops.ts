import { CommandInfo, ParamType } from "../types";

const ops: CommandInfo = {
  command: "ops",
  documentation: "Reloads and applies permissions.",
  overloads: [
    {
      params: [
        {
          value: ["list", "reload"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default ops;
