import { Parameter } from "../shared";
import { CommandInfo, ParamInfo, ParamType } from "../types";

const presetParam: ParamInfo = {
  value: [
    "minecraft:first_person",
    "minecraft:free",
    "minecraft:third_person",
    "minecraft:third_person_front",
    "minecraft:follow_orbit",
  ],
  signatureValue: "<preset>",
  type: ParamType.enum,
};

const easeParam: ParamInfo = {
  value: [
    "linear",
    "spring",
    "in_quad",
    "out_quad",
    "in_out_quad",
    "in_cubic",
    "out_cubic",
    "in_out_cubic",
    "in_quart",
    "out_quart",
    "in_out_quart",
    "in_quint",
    "out_quint",
    "in_out_quint",
    "in_sine",
    "out_sine",
    "in_out_sine",
    "in_expo",
    "out_expo",
    "in_out_expo",
    "in_circ",
    "out_circ",
    "in_out_circ",
    "in_bounce",
    "out_bounce",
    "in_out_bounce",
    "in_back",
    "out_back",
    "in_out_back",
    "in_elastic",
    "out_elastic",
    "in_out_elastic",
  ],
  signatureValue: "<easeType>",
  type: ParamType.enum,
};

const camera: CommandInfo = {
  command: "camera",
  documentation: "Issues a camera instruction",
  overloads: [
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["rot"],
          type: ParamType.keyword,
        },
        ...Parameter.rotXY,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "position",
          type: ParamType.vector3,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["rot"],
          type: ParamType.keyword,
        },
        ...Parameter.rotXY,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["ease"],
          type: ParamType.keyword,
        },
        {
          value: "<duration>",
          type: ParamType.float,
        },
        easeParam,
        {
          value: ["default"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["rot"],
          type: ParamType.keyword,
        },
        ...Parameter.rotXY,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["<target_entity>"],
          type: ParamType.entitySelector,
        },
        {
          value: "<target>",
          type: ParamType.entitySelector,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["<target_entity>"],
          type: ParamType.entitySelector,
        },
        {
          value: "<target>",
          type: ParamType.entitySelector,
        },
        {
          value: ["target_center_offset"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["remove_target"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["view_offset"],
          type: ParamType.keyword,
        },
        {
          value: "<xViewOffset>",
          type: ParamType.float,
        },
        {
          value: "<yViewOffset>",
          type: ParamType.float,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["rot"],
          type: ParamType.keyword,
        },
        ...Parameter.rotXY,
        {
          value: ["view_offset"],
          type: ParamType.keyword,
        },
        {
          value: "<xViewOffset>",
          type: ParamType.float,
        },
        {
          value: "<yViewOffset>",
          type: ParamType.float,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["pos"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["rot"],
          type: ParamType.keyword,
        },
        ...Parameter.rotXY,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["facing"],
          type: ParamType.keyword,
        },
        ...Parameter.position,
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["set"],
          type: ParamType.keyword,
        },
        presetParam,
        {
          value: ["default"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
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
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["fade"],
          type: ParamType.keyword,
        },
        {
          value: ["time"],
          type: ParamType.keyword,
        },
        {
          value: "<fadeInSeconds>",
          type: ParamType.float,
        },
        {
          value: "<holdSeconds>",
          type: ParamType.float,
        },
        {
          value: "<fadeOutSeconds>",
          type: ParamType.float,
        },
        {
          value: ["color"],
          type: ParamType.keyword,
        },
        {
          value: "<red>",
          type: ParamType.float,
        },
        {
          value: "<green>",
          type: ParamType.float,
        },
        {
          value: "<blue>",
          type: ParamType.float,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["fade"],
          type: ParamType.keyword,
        },
        {
          value: ["time"],
          type: ParamType.keyword,
        },
        {
          value: "<fadeInSeconds>",
          type: ParamType.float,
        },
        {
          value: "<holdSeconds>",
          type: ParamType.float,
        },
        {
          value: "<fadeOutSeconds>",
          type: ParamType.float,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["fade"],
          type: ParamType.keyword,
        },
        {
          value: ["color"],
          type: ParamType.keyword,
        },
        {
          value: "<red>",
          type: ParamType.float,
        },
        {
          value: "<green>",
          type: ParamType.float,
        },
        {
          value: "<blue>",
          type: ParamType.float,
        },
      ],
    },
    {
      params: [
        {
          value: "<player>",
          type: ParamType.playerSelector,
        },
        {
          value: ["fade"],
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default camera;
