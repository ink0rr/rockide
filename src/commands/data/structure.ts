import { CommandInfo, ParamType } from "../types";

const structure: CommandInfo = {
  command: "structure",
  documentation: "Saves or loads a structure in the world.",
  overloads: [
    {
      params: [
        {
          value: ["save"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
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
          value: ["disk", "memory"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["save"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
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
          value: ["true", "false"],
          type: ParamType.keyword,
        },
        {
          value: ["disk", "memory"],
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
          value: ["delete"],
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
          value: ["load"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["0_degrees", "90_degrees", "180_degrees", "270_degrees"],
          type: ParamType.keyword,
        },
        {
          value: ["x", "z", "none", "xz"],
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
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
          value: "todo",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["load"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["0_degrees", "90_degrees", "180_degrees", "270_degrees"],
          type: ParamType.keyword,
        },
        {
          value: ["x", "z", "none", "xz"],
          type: ParamType.keyword,
        },
        {
          value: ["block_by_block", "layer_by_layer"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
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
          value: "todo",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default structure;
