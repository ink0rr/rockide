import { CommandInfo, ParamType } from "../types";

const tickingarea: CommandInfo = {
  command: "tickingarea",
  documentation: "Add, remove, or list ticking areas.",
  overloads: [
    {
      params: [
        {
          value: ["add"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["add"],
          type: ParamType.keyword,
        },
        {
          value: ["circle"],
          type: ParamType.keyword,
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
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["remove"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["remove"],
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
          value: ["remove_all"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["list"],
          type: ParamType.keyword,
        },
        {
          value: ["all-dimensions"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["preload"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["preload"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default tickingarea;
