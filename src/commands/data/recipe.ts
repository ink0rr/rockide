import { CommandInfo, ParamType } from "../types";

const recipe: CommandInfo = {
  command: "recipe",
  documentation: "Unlocks recipe in the recipe book for a player.",
  overloads: [
    {
      params: [
        {
          value: ["give"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["unlockablerecipevalues", "minecraft:bundle", "*"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["take"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["unlockablerecipevalues", "minecraft:bundle", "*"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default recipe;
