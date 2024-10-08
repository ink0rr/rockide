import { CommandInfo, ParamType } from "../types";

const me: CommandInfo = {
  command: "me",
  documentation: "Displays a message about yourself.",
  overloads: [
    {
      params: [
        {
          value: "string",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default me;
