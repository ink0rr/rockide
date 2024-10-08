import * as vscode from "vscode";
import { RockideContext } from "../context";
import { Rockide } from "../rockide";

export type RockideHandler = {
  pattern: string | string[];
  index?: "parse" | "path" | "mcfunction";
  process?(ctx: RockideContext, rockide: Rockide): RockideProcess | void;
};

export type RockideProcess = {
  completions?(): Array<string | vscode.CompletionItem>;
  definitions?(): Promise<vscode.LocationLink>[] | vscode.LocationLink[] | void;
  signature?(): vscode.SignatureHelp | undefined;
};
