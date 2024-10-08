import { CommandInfo, ParamType } from "../types";

const xp: CommandInfo = {
  command: "xp",
  documentation: "Adds or removes player experience.",
  overloads: [
    {
      params: [
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
      ],
    },
    {
      params: [
        {
          value: ["20", "20L"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
      ],
    },
  ],
};
export default xp;
