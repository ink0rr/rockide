import { CommandInfo, ParamType } from "../types";

const weather: CommandInfo = {
  command: "weather",
  documentation: "Sets the weather.",
  overloads: [
    {
      params: [
        {
          value: ["clear", "rain", "thunder"],
          type: ParamType.keyword,
        },
        {
          value: "int",
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
      ],
    },
  ],
};
export default weather;
