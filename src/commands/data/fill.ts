import { CommandInfo, ParamType } from "../types";

const fill: CommandInfo = {
  command: "fill",
  documentation: "Fills all or parts of a region with a specific block.",
  overloads: [
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "blockState",
          type: ParamType.keyword,
        },
        {
          value: ["outline", "hollow", "destroy", "keep"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: ["outline", "hollow", "destroy", "keep"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "blockState",
          type: ParamType.keyword,
        },
        {
          value: ["replace"],
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "blockState",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: ["replace"],
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "blockState",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default fill;
