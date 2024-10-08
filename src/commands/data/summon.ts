import { CommandInfo, ParamType } from "../types";

const summon: CommandInfo = {
  command: "summon",
  documentation: "Summons an entity.",
  overloads: [
    {
      params: [
        {
          value: ["entity_type", "minecraft:slime"],
          type: ParamType.keyword,
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
          value: ["entity_events", "minecraft:on_spawned"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["entity_type", "minecraft:slime"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["entity_type", "minecraft:slime"],
          type: ParamType.keyword,
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
          value: ["entity_events", "minecraft:on_spawned"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["entity_type", "minecraft:slime"],
          type: ParamType.keyword,
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
          value: ["entity_events", "minecraft:on_spawned"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default summon;
