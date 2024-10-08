import { CommandInfo, ParamType } from "../types";

const give: CommandInfo = {
  command: "give",
  documentation: "Gives an item to a player.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_JSON_OBJECT"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default give;
