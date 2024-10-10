import { enchantments } from "../../literals/enchantments";
import { CommandInfo, ParamType } from "../types";

const enchant: CommandInfo = {
  command: "enchant",
  documentation: "Adds an enchantment to a player's selected item.",
  overloads: [
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: enchantments,
          signatureValue: "<enchantmentName>",
          type: ParamType.enum,
        },
        {
          value: "[level]",
          type: ParamType.number,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["0"],
          documentation: "By ID",
          signatureValue: "<enchantmentId>",
          type: ParamType.enum,
        },
        {
          value: "[level]",
          type: ParamType.number,
        },
      ],
    },
  ],
};
export default enchant;
