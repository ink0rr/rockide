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
          value: "<npc>",
          type: ParamType.entitySelector,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: "<sceneName>",
          type: ParamType.string,
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
          value: "<npc>",
          type: ParamType.entitySelector,
        },
        {
          value: "<sceneName>",
          type: ParamType.string,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
      ],
    },
  ],
};
export default dialogue;
