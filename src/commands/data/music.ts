import { CommandInfo, ParamType } from "../types";

const music: CommandInfo = {
  command: "music",
  documentation: "Allows you to control playing music tracks.",
  overloads: [
    {
      params: [
        {
          value: ["queue"],
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
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["play_once", "loop"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["play"],
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
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["play_once", "loop"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["stop"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["volume"],
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
export default music;
