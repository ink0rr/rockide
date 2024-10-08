import { CommandInfo, ParamType } from "../types";

const scriptevent: CommandInfo = {
  command: "scriptevent",
  documentation: "Triggers a script event with an ID and message.",
  overloads: [
    {
      params: [
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "string",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default scriptevent;
