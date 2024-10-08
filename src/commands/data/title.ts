import { CommandInfo, ParamType } from "../types";

const title: CommandInfo = {
  command: "title",
  documentation: "Controls screen titles.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["clear"],
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
          value: ["reset"],
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
          value: ["title", "subtitle", "actionbar"],
          type: ParamType.keyword,
        },
        {
          value: "string",
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
          value: ["times"],
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
export default title;
