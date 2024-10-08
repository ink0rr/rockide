import { CommandInfo, ParamType } from "../types";

const alwaysday: CommandInfo = {
  command: "alwaysday",
  documentation: "Locks and unlocks the day-night cycle.",
  overloads: [
    {
      params: [
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default alwaysday;
