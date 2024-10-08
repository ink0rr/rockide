import { CommandInfo, ParamType } from "../types";

const ride: CommandInfo = {
  command: "ride",
  documentation:
    "Makes entities ride other entities, stops entities from riding, makes rides evict their riders, or summons rides or riders.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["start_riding"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["teleport_rider", "teleport_ride"],
          type: ParamType.keyword,
        },
        {
          value: ["until_full", "if_group_fits"],
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
          value: ["stop_riding"],
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
          value: ["evict_riders"],
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
          value: ["summon_rider"],
          type: ParamType.keyword,
        },
        {
          value: ["entity_type", "minecraft:slime"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
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
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["summon_ride"],
          type: ParamType.keyword,
        },
        {
          value: ["entity_type", "minecraft:slime"],
          type: ParamType.keyword,
        },
        {
          value: ["no_ride_change", "reassign_rides", "skip_riders"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
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
export default ride;
