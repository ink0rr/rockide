import { spellEffects } from "../../literals/spell_effects";
import { CommandInfo, ParamType } from "../types";

const effect: CommandInfo = {
  command: "effect",
  documentation: "Add or remove status effects.",
  overloads: [
    {
      params: [
        {
          value: "<target>",
          type: ParamType.entitySelector,
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
          value: "<target>",
          type: ParamType.entitySelector,
        },
        {
          value: spellEffects,
          signatureValue: "<effect>",
          type: ParamType.enum,
        },
        {
          value: "[seconds]",
          type: ParamType.number,
        },
        {
          value: "[amplifier]",
          type: ParamType.number,
        },
        {
          value: ["true", "false"],
          signatureValue: "[hideParticles]",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default effect;
