import { CommandInfo, ParamType } from "../types";

const schedule: CommandInfo = {
  command: "schedule",
  documentation: "Schedules an action to be executed once an area is loaded, or after a certain amount of time.",
  overloads: [
    {
      params: [
        {
          value: ["on_area_loaded"],
          type: ParamType.keyword,
        },
        {
          value: ["add"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "mcfunctionPath",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["on_area_loaded"],
          type: ParamType.keyword,
        },
        {
          value: ["add"],
          type: ParamType.keyword,
        },
        {
          value: ["circle"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "mcfunctionPath",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["on_area_loaded"],
          type: ParamType.keyword,
        },
        {
          value: ["add"],
          type: ParamType.keyword,
        },
        {
          value: ["tickingarea"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "mcfunctionPath",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default schedule;
