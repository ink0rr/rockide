import * as JSONC from "jsonc-parser";
import * as vscode from "vscode";

export const baseGlob = "**/!(build|dist|tmp)";

export const bpGlob = "{behavior_pack,*BP,BP_*,*bp,bp_*}";

export const rpGlob = "{resource_pack,*RP,RP_*,*rp,rp_*}";

export const projectGlob = "{behavior_pack,*BP,BP_*,*bp,bp_*,resource_pack,*RP,RP_*,*rp,rp_*}";

export const NullNode: JSONC.Node = {
  type: "null",
  length: 4,
  offset: 0,
};

export const jsonSelector: vscode.DocumentSelector = [
  { scheme: "file", language: "json" },
  { scheme: "file", language: "jsonc" },
];

export const commandSelector: vscode.DocumentSelector = jsonSelector.concat({
  scheme: "file",
  language: "rockide-mcfunction",
});
