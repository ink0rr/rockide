import { CommandInfo, ParamType } from "../types";

const save: CommandInfo = {
  command: "save",
  documentation: "Control or check how the game saves data to disk.",
  overloads: [
    {
      params: [
        {
          value: ["query", "hold", "resume"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default save;
