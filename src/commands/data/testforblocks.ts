import { CommandInfo, ParamType } from "../types";

const testforblocks: CommandInfo = {
  command: "testforblocks",
  documentation: "Tests whether the blocks in two regions match.",
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
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["masked", "all"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default testforblocks;
