import { CommandInfo, ParamType } from "../types";

const setblock: CommandInfo = {
  command: "setblock",
  documentation: "Changes a block to another block.",
  overloads: [
    {
      params: [
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
          value: ["replace", "destroy", "keep"],
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
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: ["replace", "destroy", "keep"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default setblock;
