import { CommandInfo, ParamType } from "../types";

const replaceitem: CommandInfo = {
  command: "replaceitem",
  documentation: "Replaces items in inventories.",
  overloads: [
    {
      params: [
        {
          value: ["block"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["slot.container"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_JSON_OBJECT"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["entity"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: [
            "slot.weapon.mainhand",
            "slot.weapon.offhand",
            "slot.armor.head",
            "slot.armor.chest",
            "slot.armor.legs",
            "slot.armor.feet",
            "slot.armor.body",
            "slot.hotbar",
            "slot.inventory",
            "slot.enderchest",
            "slot.saddle",
            "slot.armor",
            "slot.chest",
            "slot.equippable",
          ],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_JSON_OBJECT"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["block"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["slot.container"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["destroy", "keep"],
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_JSON_OBJECT"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["entity"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: [
            "slot.weapon.mainhand",
            "slot.weapon.offhand",
            "slot.armor.head",
            "slot.armor.chest",
            "slot.armor.legs",
            "slot.armor.feet",
            "slot.armor.body",
            "slot.hotbar",
            "slot.inventory",
            "slot.enderchest",
            "slot.saddle",
            "slot.armor",
            "slot.chest",
            "slot.equippable",
          ],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["destroy", "keep"],
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["unknown_JSON_OBJECT"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default replaceitem;
