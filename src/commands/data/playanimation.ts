import { CommandInfo, ParamType } from "../types";

const playanimation: CommandInfo = {
  command: "playanimation",
  documentation: "Makes one or more entities play a one-off animation. Assumes all variables are setup correctly.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default playanimation;
