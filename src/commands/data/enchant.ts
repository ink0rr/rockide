import { CommandInfo, ParamType } from "../types";

const enchant: CommandInfo = {
  command: "enchant",
  documentation: "Adds an enchantment to a player's selected item.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: [
            "protection",
            "fire_protection",
            "feather_falling",
            "blast_protection",
            "projectile_protection",
            "thorns",
            "respiration",
            "depth_strider",
            "aqua_affinity",
            "sharpness",
            "smite",
            "bane_of_arthropods",
            "knockback",
            "fire_aspect",
            "looting",
            "efficiency",
            "silk_touch",
            "unbreaking",
            "fortune",
            "power",
            "punch",
            "flame",
            "infinity",
            "luck_of_the_sea",
            "lure",
            "frost_walker",
            "mending",
            "binding",
            "vanishing",
            "impaling",
            "riptide",
            "loyalty",
            "channeling",
            "multishot",
            "piercing",
            "quick_charge",
            "soul_speed",
            "swift_sneak",
            "wind_burst",
            "density",
            "breach",
          ],
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
          value: "selector",
          type: ParamType.playerSelector,
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
  ],
};
export default enchant;
