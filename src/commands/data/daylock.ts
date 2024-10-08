import { CommandInfo, ParamType } from "../types";

const daylock: CommandInfo = {
  command: "daylock",
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
export default daylock;
