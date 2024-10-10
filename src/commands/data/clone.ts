import { blockIdentifier } from "../../literals/block_identifier";
import { Parameter } from "../parameter";
import { CommandInfo, ParamType } from "../types";

const clone: CommandInfo = {
  command: "clone",
  documentation: "Clones blocks from one region to another.",
  overloads: [
    {
      params: [
        ...Parameter.position,
        ...Parameter.position,
        ...Parameter.position,
        {
          value: ["replace", "masked"],
          type: ParamType.keyword,
        },
        {
          value: ["normal", "force", "move"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        ...Parameter.position,
        ...Parameter.position,
        ...Parameter.position,
        {
          value: ["filtered"],
          type: ParamType.keyword,
        },
        {
          value: ["normal", "force", "move"],
          type: ParamType.keyword,
        },
        {
          // todo: defined blocks
          value: blockIdentifier,
          type: ParamType.keyword,
        },
        {
          // todo: handle this
          value: "blockState",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default clone;
