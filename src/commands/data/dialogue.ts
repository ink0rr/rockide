import { CommandInfo, ParamType } from "../types";

const dialogue: CommandInfo = {
  command: "dialogue",
  documentation: "Opens NPC dialogue for a player.",
  overloads: [
    {
      params: [
        {
          value: ["open"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
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
    {
      params: [
        {
          value: ["change"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
      ],
    },
  ],
};
export default dialogue;
