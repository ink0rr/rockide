import { CommandInfo, ParamType } from "../types";

const camerashake: CommandInfo = {
  command: "camerashake",
  documentation: "Applies shaking to the players' camera with a specified intensity and duration.",
  overloads: [
    {
      params: [
        {
          value: ["add"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: "<intensity>",
          type: ParamType.float,
        },
        {
          value: "<seconds>",
          type: ParamType.float,
        },
        {
          value: ["positional", "rotational"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["stop"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
      ],
    },
  ],
};
export default camerashake;
