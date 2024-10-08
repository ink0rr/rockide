import { CommandInfo, ParamType } from "../types";

const inputpermission: CommandInfo = {
  command: "inputpermission",
  documentation: "Sets whether or not a player's input can affect their character.",
  overloads: [
    {
      params: [
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["movement", "camera"],
          type: ParamType.keyword,
        },
        {
          value: ["enabled", "disabled"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["query"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["movement", "camera"],
          type: ParamType.keyword,
        },
        {
          value: ["enabled", "disabled"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default inputpermission;
