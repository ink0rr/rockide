import { CommandInfo, ParamType } from "../types";

const particle: CommandInfo = {
  command: "particle",
  documentation: "Creates a particle emitter",
  overloads: [
    {
      params: [
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default particle;
