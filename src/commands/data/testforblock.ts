import { CommandInfo, ParamType } from "../types";

const testforblock: CommandInfo = {
  command: "testforblock",
  documentation: "Tests whether a certain block is in a specific location.",
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
      ],
    },
  ],
};
export default testforblock;
