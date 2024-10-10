import { ParamInfo, ParamType } from "./types";

export namespace Parameter {
  export const position: ParamInfo[] = [
    {
      value: "<x>",
      type: ParamType.location,
    },
    {
      value: "<y>",
      type: ParamType.location,
    },
    {
      value: "<z>",
      type: ParamType.location,
    },
  ];
  export const rotXY: ParamInfo[] = [
    {
      value: "<pitch>",
      type: ParamType.pitch,
    },
    {
      value: "<yaw>",
      type: ParamType.yaw,
    },
  ];
  export const rotYX: ParamInfo[] = [
    {
      value: "<yaw>",
      type: ParamType.yaw,
    },
    {
      value: "<pitch>",
      type: ParamType.pitch,
    },
  ];
}
