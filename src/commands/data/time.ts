import { CommandInfo, ParamType } from "../types";

const time: CommandInfo = {
  command: "time",
  documentation: "Changes or queries the world's game time.",
  overloads: [
    {
      params: [
        {
          value: ["add"],
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
          value: ["set"],
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
          value: ["set"],
          type: ParamType.keyword,
        },
        {
          value: ["day", "sunrise", "noon", "sunset", "night", "midnight"],
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
          value: ["daytime", "gametime", "day"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default time;
