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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "rot",
          type: ParamType.keyword,
        },
        {
          value: "rot",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["eyes", "feet"],
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "blockState",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "selector",
          type: ParamType.playerSelector,
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: [],
          type: ParamType.keyword,
        },
        {
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: "selector",
          type: ParamType.playerSelector,
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
          value: ["unknown_EXECUTECHAINEDOPTION_0"],
          type: ParamType.keyword,
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
          value: ["unknown_CODEBUILDERARGS"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default execute;
