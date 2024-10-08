import { CommandInfo, ParamType } from "../types";

const functionCommand: CommandInfo = {
  command: "function",
  documentation: "Runs commands found in the corresponding function file.",
  overloads: [
    {
      params: [
        {
          value: "mcfunctionPath",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default functionCommand;
