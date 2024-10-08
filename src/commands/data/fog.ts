import { CommandInfo, ParamType } from "../types";

const fog: CommandInfo = {
  command: "fog",
  documentation: "Add or remove fog settings file",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["push"],
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
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["pop", "remove"],
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
export default fog;
