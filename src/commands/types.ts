import { MarkdownString } from "vscode";

export type CommandInfo = {
  command: string;
  overloads?: Overload[];
  documentation?: string | MarkdownString;
};

export type Overload = {
  params: ParamInfo[];
  documentation?: string | MarkdownString;
};

export type ParamInfo = {
  value: string | string[];
  signatureValue?: string;
  type: ParamType;
  required?: boolean;
  documentation?: string | MarkdownString;
};

export enum ParamType {
  keyword = "keyword",
  enum = "enum",
  playerSelector = "playerSelector",
  entitySelector = "entitySelector",
  selectorWildcard = "selectorWildcard",
  location = "location",
  tag = "tag",
  string = "string",
  number = "number",
  float = "float",
  vector3 = "vector3",
  yaw = "yaw",
  pitch = "pitch",
}
