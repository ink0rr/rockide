import { itemIdentifier } from "../../literals/item_identifier";
import { CommandInfo, ParamType } from "../types";

const clear: CommandInfo = {
  command: "clear",
  documentation: "Clears items from player inventory.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          // todo: defined items
          value: itemIdentifier,
          signatureValue: "<itemIdentifier>",
          documentation: "The item to be cleared.",
          type: ParamType.keyword,
        },
        {
          value: "<data>",
          documentation: "The data value of the item to be cleared.",
          type: ParamType.number,
        },
        {
          value: "<maxCount>",
          documentation: "The maximum number of items to be cleared.",
          type: ParamType.number,
        },
      ],
    },
  ],
};
export default clear;
