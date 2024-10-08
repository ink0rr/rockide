import { CommandInfo, ParamType } from "../types";

const difficulty: CommandInfo = {
  command: "difficulty",
  documentation: "Sets the difficulty level.",
  overloads: [
    {
      params: [
        {
          value: ["normal", "peaceful", "easy", "hard", "p", "e", "n", "h"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "int",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default difficulty;
