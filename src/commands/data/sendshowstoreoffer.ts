import { CommandInfo, ParamType } from "../types";

const sendshowstoreoffer: CommandInfo = {
  command: "sendshowstoreoffer",
  documentation: "Sends a request to show a store offer to the target player.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["marketplace", "character"],
          type: ParamType.keyword,
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["server"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default sendshowstoreoffer;
