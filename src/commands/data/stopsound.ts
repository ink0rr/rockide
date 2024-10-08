import { CommandInfo, ParamType } from "../types";

const stopsound: CommandInfo = {
  command: "stopsound",
  documentation: "Stops a sound.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default stopsound;
