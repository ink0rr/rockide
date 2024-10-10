import { CommandInfo, ParamType } from "../types";

const difficulty: CommandInfo = {
  command: "difficulty",
  documentation: "Sets the difficulty level.",
  overloads: [
    {
      params: [
        {
          value: ["normal", "peaceful", "easy", "hard", "p", "e", "n", "h"],
          signatureValue: "<difficulty>",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["0", "1", "2", "3"],
          // documentation: "0: peaceful, 1: easy, 2: normal, 3: hard",
          documentation: ["peaceful", "easy", "normal", "hard"],
          signatureValue: "<difficulty>",
          type: ParamType.enum,
        },
      ],
    },
  ],
};
export default difficulty;
