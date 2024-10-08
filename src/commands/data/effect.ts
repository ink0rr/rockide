import { CommandInfo, ParamType } from "../types";

const effect: CommandInfo = {
  command: "effect",
  documentation: "Add or remove status effects.",
  overloads: [
    {
      params: [
        {
          value: "selector",
          type: ParamType.playerSelector,
        },
        {
          value: ["clear"],
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
          value: [
            "wither",
            "speed",
            "slowness",
            "haste",
            "mining_fatigue",
            "strength",
            "instant_health",
            "instant_damage",
            "jump_boost",
            "nausea",
            "regeneration",
            "resistance",
            "fire_resistance",
            "water_breathing",
            "invisibility",
            "blindness",
            "night_vision",
            "hunger",
            "weakness",
            "poison",
            "health_boost",
            "absorption",
            "saturation",
            "levitation",
            "fatal_poison",
            "conduit_power",
            "slow_falling",
            "bad_omen",
            "village_hero",
            "darkness",
            "trial_omen",
            "wind_charged",
            "weaving",
            "oozing",
            "infested",
            "raid_omen",
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
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default effect;
