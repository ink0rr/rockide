import { CommandInfo, ParamType } from "../types";

const script: CommandInfo = {
  command: "script",
  documentation: "Debugging options for GameTest Framework.",
  overloads: [
    {
      params: [
        {
          value: ["profiler"],
          type: ParamType.keyword,
        },
        {
          value: ["start"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["profiler"],
          type: ParamType.keyword,
        },
        {
          value: ["stop"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["watchdog"],
          type: ParamType.keyword,
        },
        {
          value: ["exportstats"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default script;
