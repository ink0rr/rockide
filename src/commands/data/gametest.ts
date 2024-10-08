import { CommandInfo, ParamType } from "../types";

const gametest: CommandInfo = {
  command: "gametest",
  documentation: "Interacts with gametest.",
  overloads: [
    {
      params: [
        {
          value: ["runthis"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["run"],
          type: ParamType.keyword,
        },
        {
          value: [],
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
          value: ["run"],
          type: ParamType.keyword,
        },
        {
          value: [],
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
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
    {
      params: [
        {
          value: ["runset"],
          type: ParamType.keyword,
        },
        {
          value: [],
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
          value: ["runsetuntilfail"],
          type: ParamType.keyword,
        },
        {
          value: [],
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
          value: ["clearall"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["create"],
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
          value: "int",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["runthese"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["stopall"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default gametest;
