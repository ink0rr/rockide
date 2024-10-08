import { CommandInfo, ParamType } from "../types";

const tellraw: CommandInfo = {
  command: "tellraw",
  documentation: "Sends a JSON message to players.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["unknown_JSON_OBJECT"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default tellraw;
