import { CommandInfo, ParamType } from "../types";

const setworldspawn: CommandInfo = {
  command: "setworldspawn",
  documentation: "Sets the world spawn.",
  overloads: [
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default setworldspawn;
