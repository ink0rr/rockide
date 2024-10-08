import * as vscode from "vscode";
import { jsonSelector } from "./constants";
import { JsonProvider } from "./providers/json";
import { MolangProvider } from "./providers/molang";
import { Rockide } from "./rockide";
import { legend } from "./semantics";

export async function activate(context: vscode.ExtensionContext) {
  const rockide = new Rockide();
  const isBedrock = await rockide.checkWorkspace();
  if (!isBedrock) {
    console.log("Not a Bedrock Workspace, Rockide is disabled.");
    return;
  }
  console.log("Rockide activated!");
  await rockide.indexWorkspace();

  const jsonProvider = new JsonProvider(rockide);
  const molangProvider = new MolangProvider();

  context.subscriptions.push(
    vscode.commands.registerCommand("rockide.reloadWorkspace", () => rockide.indexWorkspace()),

    // json
    vscode.languages.registerCompletionItemProvider(jsonSelector, jsonProvider, "."),
    vscode.languages.registerDefinitionProvider(jsonSelector, jsonProvider),
    vscode.workspace.onDidChangeTextDocument((e) => jsonProvider.onDidChangeTextDocument(e)),
    vscode.workspace.onDidCreateFiles((e) => jsonProvider.onDidCreateFiles(e)),
    vscode.workspace.onDidRenameFiles((e) => jsonProvider.onDidRenameFiles(e)),
    vscode.workspace.onDidDeleteFiles((e) => jsonProvider.onDidDeleteFiles(e)),

    // molang
    vscode.languages.registerDocumentSemanticTokensProvider(jsonSelector, molangProvider, legend),
  );
}

export function deactivate() {}
