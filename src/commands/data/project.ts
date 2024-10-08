import { CommandInfo, ParamType } from "../types";

const project: CommandInfo = {
  command: "project",
  documentation: "Manipulate the currently loaded project",
  overloads: [
    {
      params: [
        {
          value: ["export"],
          type: ParamType.keyword,
        },
        {
          value: ["project", "template", "world"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default project;
