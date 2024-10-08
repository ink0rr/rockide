import { CommandInfo, ParamType } from "../types";

const tag: CommandInfo = {
  command: "tag",
  documentation: "Manages tags stored in entities.",
  overloads: [
    {
      params: [
        {
          value: "<player>",
          type: ParamType.selectorWildcard,
        },
        {
          value: ["add", "remove"],
          type: ParamType.keyword,
        },
        {
          value: "<name>",
          type: ParamType.string,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.selectorWildcard,
        },
        {
          value: ["list"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default tag;
