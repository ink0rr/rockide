import { randomUUID } from "crypto";
import * as vscode from "vscode";
import { projectGlob } from "../../../constants";
import { JsonHandler } from "./_type";

export const manifest: JsonHandler = {
  pattern: `**/${projectGlob}/manifest.json`,
  index: "parse",
  process(ctx, rockide) {
    if (ctx.matchArrayObject("dependencies", "uuid")) {
      return {
        completions: () =>
          rockide
            .getManifests()
            .filter(({ path }) => path !== ctx.document.uri.fsPath)
            .flatMap(({ values }) => values),
        definitions: () =>
          rockide
            .getManifests()
            .filter(({ values }) => values.includes(ctx.nodeValue))
            .map(({ path, root }) => ctx.createDefinition(path, root)),
      };
    }
    if (ctx.matchField("uuid")) {
      return {
        completions: () => {
          const generateUuid = new vscode.CompletionItem("Generate UUID", vscode.CompletionItemKind.Constant);
          generateUuid.insertText = randomUUID();
          return [generateUuid];
        },
      };
    }
  },
};
