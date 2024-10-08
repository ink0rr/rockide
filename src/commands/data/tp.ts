import { CommandInfo, ParamType } from "../types";

const tp: CommandInfo = {
  command: "tp",
  documentation: "Teleports entities (players, mobs, etc.).",
  overloads: [
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "position",
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
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "position",
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
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
      ],
    },
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default tp;
