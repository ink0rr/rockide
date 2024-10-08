import { damageSource } from "../../literals/damage_source";
import { CommandInfo, ParamType } from "../types";

const damage: CommandInfo = {
  command: "damage",
  documentation: "Apply damage to the specified entities.",
  overloads: [
    {
      params: [
        {
          value: "<target>",
          type: ParamType.entitySelector,
        },
        {
          value: "<amount>",
          type: ParamType.number,
        },
        {
          value: damageSource,
          signatureValue: "<cause>",
          type: ParamType.enum,
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
          value: "<amount>",
          type: ParamType.number,
        },
        {
          value: damageSource,
          signatureValue: "<cause>",
          type: ParamType.enum,
        },
        {
          value: ["entity"],
          type: ParamType.keyword,
        },
        {
          value: "<target>",
          type: ParamType.entitySelector,
        },
      ],
    },
  ],
};
export default damage;
