import { CommandInfo, ParamType } from "../types";

const loot: CommandInfo = {
  command: "loot",
  documentation: "Drops the given loot table into the specified inventory or into the world.",
  overloads: [
    {
      params: [
        {
          value: ["spawn"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["loot"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["spawn"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["kill"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["give"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["loot"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["give"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["kill"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["insert"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["loot"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["insert"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.keyword,
        },
        {
          value: ["kill"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["loot"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: ["loot"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["kill"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: ["kill"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: ["slot.container"],
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
          value: ["loot"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: ["slot.container"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["loot"],
          type: ParamType.keyword,
        },
        {
          value: "todo",
          type: ParamType.keyword,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: ["slot.container"],
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
          value: ["kill"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: ["replace"],
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
          value: ["slot.container"],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
        {
          value: ["kill"],
          type: ParamType.keyword,
        },
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["minecraft:dirt", "dirt", "todo"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default loot;
