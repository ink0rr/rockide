import { blockIdentifier } from "../../literals/block_identifier";
import { Parameter } from "../parameter";
import { CommandInfo, ParamType } from "../types";

const execute: CommandInfo = {
  command: "execute",
  documentation: "Executes a command on behalf of one or more entities.",
  overloads: [
    {
      params: [
        {
          value: ["as"],
          type: ParamType.keyword,
        },
        {
          value: "<origin>",
          type: ParamType.entitySelector,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["at"],
          type: ParamType.keyword,
        },
        {
          value: "<origin>",
          type: ParamType.entitySelector,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["in"],
          type: ParamType.keyword,
        },
        {
          value: ["overworld", "nether", "the_end"],
          signatureValue: "<dimension>",
          type: ParamType.keyword,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["positioned"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["positioned"],
          type: ParamType.keyword,
        },
        {
          value: ["as"],
          type: ParamType.keyword,
        },
        {
          value: "<origin>",
          type: ParamType.entitySelector,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["rotated"],
          type: ParamType.keyword,
        },
        ...Parameter.rotYX,
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["rotated"],
          type: ParamType.keyword,
        },
        {
          value: ["as"],
          type: ParamType.keyword,
        },
        {
          value: "<origin>",
          type: ParamType.entitySelector,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: ["entity"],
          type: ParamType.keyword,
        },
        {
          value: "<origin>",
          type: ParamType.entitySelector,
        },
        {
          value: ["eyes", "feet"],
          type: ParamType.keyword,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["align"],
          type: ParamType.keyword,
        },
        {
          value: ["x", "y", "z", "xy", "xz", "yz", "xyz"],
          type: ParamType.enum,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["anchored"],
          type: ParamType.keyword,
        },
        {
          value: ["eyes", "feet"],
          type: ParamType.keyword,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["if", "unless"],
          type: ParamType.keyword,
        },
        {
          value: ["block"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: blockIdentifier,
          type: ParamType.enum,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["if", "unless"],
          type: ParamType.keyword,
        },
        {
          value: ["block"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: blockIdentifier,
          type: ParamType.enum,
        },
        {
          value: "blockState",
          type: ParamType.keyword,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["if", "unless"],
          type: ParamType.keyword,
        },
        {
          value: ["blocks"],
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
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["masked", "all"],
          type: ParamType.keyword,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["if", "unless"],
          type: ParamType.keyword,
        },
        {
          value: ["entity"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["if", "unless"],
          type: ParamType.keyword,
        },
        {
          value: ["score"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.selectorWildcard,
        },
        {
          value: [],
          type: ParamType.keyword,
        },
        {
          value: ["%=", "*=", "+=", "-=", "/=", "<", "=", ">", "><", "matches"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.selectorWildcard,
        },
        {
          value: [],
          type: ParamType.keyword,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["if", "unless"],
          type: ParamType.keyword,
        },
        {
          value: ["score"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.selectorWildcard,
        },
        {
          value: [],
          type: ParamType.keyword,
        },
        {
          value: ["matches"],
          type: ParamType.keyword,
        },
        {
          value: "0..10",
          type: ParamType.keyword,
        },
        {
          value: ["<chainedCommand>"],
          type: ParamType.executeChainedOption,
        },
      ],
    },
    {
      params: [
        {
          value: ["run"],
          type: ParamType.keyword,
        },
        {
          value: ["<subcommand>"],
          type: ParamType.subcommnad,
        },
      ],
    },
  ],
};
export default execute;
