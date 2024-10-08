import { CommandInfo, ParamType } from "../types";

const spreadplayers: CommandInfo = {
  command: "spreadplayers",
  documentation: "Teleports entities to random locations.",
  overloads: [
    {
      params: [
        {
          value: "rot",
          type: ParamType.keyword,
        },
        {
          value: "rot",
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "rot",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default spreadplayers;
