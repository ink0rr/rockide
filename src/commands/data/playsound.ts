import { CommandInfo, ParamType } from "../types";

const playsound: CommandInfo = {
  command: "playsound",
  documentation: "Plays a sound.",
  overloads: [
    {
      params: [
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "int",
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
export default playsound;
