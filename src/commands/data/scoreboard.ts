import { CommandInfo, ParamType } from "../types";

const scoreboard: CommandInfo = {
  command: "scoreboard",
  documentation: "Tracks and displays scores for various objectives.",
  overloads: [
    {
      params: [
        {
          value: ["objectives"],
          type: ParamType.keyword,
          documentation: "Manages objectives.",
        },
        {
          value: ["add"],
          type: ParamType.keyword,
          documentation: "Adds a new objective to the scoreboard.",
        },
        {
          value: "<name>",
          type: ParamType.string,
          documentation: "The unique name of the objective.",
        },
        {
          value: ["dummy"],
          type: ParamType.keyword,
          documentation: "The type of the objective.",
        },
        {
          value: "[displayName]",
          type: ParamType.string,
          documentation: "The display name of the objective.",
        },
      ],
    },
    {
      params: [
        {
          value: ["objectives"],
          type: ParamType.keyword,
          documentation: "Manages objectives.",
        },
        {
          value: ["remove"],
          type: ParamType.keyword,
          documentation: "Removes an objective from the scoreboard.",
        },
        {
          value: "<name>",
          type: ParamType.string,
          documentation: "The unique name of the objective.",
        },
      ],
    },
    {
      params: [
        {
          value: ["objectives"],
          type: ParamType.keyword,
          documentation: "Manages objectives.",
        },
        {
          value: ["list"],
          type: ParamType.keyword,
          documentation: "Lists all objectives on the scoreboard.",
        },
      ],
    },
    {
      params: [
        {
          value: ["objectives"],
          type: ParamType.keyword,
          documentation: "Manages objectives.",
        },
        {
          value: ["setdisplay"],
          type: ParamType.keyword,
          documentation: "Sets the display of an objective.",
        },
        {
          value: ["list", "sidebar"],
          type: ParamType.keyword,
          documentation: "The display type of the objective.",
        },
        {
          value: "<name>",
          type: ParamType.string,
          documentation: "The unique name of the objective.",
        },
        {
          value: ["ascending", "descending"],
          type: ParamType.keyword,
          documentation: "The order of the objective.",
        },
      ],
    },
    {
      params: [
        {
          value: ["objectives"],
          type: ParamType.keyword,
          documentation: "Manages objectives.",
        },
        {
          value: ["setdisplay"],
          type: ParamType.keyword,
          documentation: "Sets the display of an objective.",
        },
        {
          value: ["belowname"],
          type: ParamType.keyword,
          documentation: "The display type of the objective.",
        },
        {
          value: "<name>",
          type: ParamType.string,
          documentation: "The unique name of the objective.",
        },
      ],
    },
    {
      params: [
        {
          value: ["players"],
          type: ParamType.keyword,
          documentation: "Manages player scores.",
        },
        {
          value: ["list"],
          type: ParamType.keyword,
          documentation: "Lists all scores for the player.",
        },
        {
          value: "<selector>",
          type: ParamType.selectorWildcard,
          documentation: "The player to list the scores for.",
        },
      ],
    },
    {
      params: [
        {
          value: ["players"],
          type: ParamType.keyword,
        },
        {
          value: ["reset"],
          type: ParamType.keyword,
        },
        {
          value: "<selector>",
          type: ParamType.selectorWildcard,
        },
        {
          value: "<name>",
          type: ParamType.string,
        },
      ],
    },
    {
      params: [
        {
          value: ["players"],
          type: ParamType.keyword,
        },
        {
          value: ["test"],
          type: ParamType.keyword,
        },
        {
          value: "<selector>",
          type: ParamType.selectorWildcard,
        },
        {
          value: "<name>",
          type: ParamType.string,
        },
        {
          value: ["0", "1", "2"],
          type: ParamType.keyword,
        },
        {
          value: ["0", "1", "2"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["players"],
          type: ParamType.keyword,
        },
        {
          value: ["random"],
          type: ParamType.keyword,
        },
        {
          value: "<selector>",
          type: ParamType.selectorWildcard,
        },
        {
          value: "<name>",
          type: ParamType.string,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["players"],
          type: ParamType.keyword,
        },
        {
          value: ["set", "add", "remove"],
          type: ParamType.keyword,
        },
        {
          value: "<selector>",
          type: ParamType.selectorWildcard,
        },
        {
          value: "<name>",
          type: ParamType.string,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["players"],
          type: ParamType.keyword,
        },
        {
          value: ["operation"],
          type: ParamType.keyword,
        },
        {
          value: "<selector>",
          type: ParamType.selectorWildcard,
        },
        {
          value: "<name>",
          type: ParamType.string,
        },
        {
          value: ["%=", "*=", "+=", "-=", "/=", "<", "=", ">", "><"],
          type: ParamType.keyword,
        },
        {
          value: "<selector>",
          type: ParamType.selectorWildcard,
        },
        {
          value: "<name>",
          type: ParamType.string,
        },
      ],
    },
  ],
};
export default scoreboard;
