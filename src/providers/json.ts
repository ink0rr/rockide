import * as JSONC from "jsonc-parser";
import { isMatch } from "micromatch";
import * as vscode from "vscode";
import { baseGlob, projectGlob } from "../constants";
import { createContext } from "../context";
import { fileHandlers } from "../handlers";
import { Rockide } from "../rockide";

export class JsonProvider implements vscode.CompletionItemProvider, vscode.DefinitionProvider {
  constructor(private rockide: Rockide) {}
  provideCompletionItems(
    document: vscode.TextDocument,
    position: vscode.Position,
  ): vscode.ProviderResult<vscode.CompletionItem[] | vscode.CompletionList<vscode.CompletionItem>> {
    for (const handler of fileHandlers) {
      if (isMatch(document.uri.fsPath, handler.pattern)) {
        if (handler.process) {
          const ctx = createContext(document, position);
          const completions = handler.process(ctx, this.rockide)?.completions?.();
          return completions?.map((value) => ctx.createCompletion(value));
        }
      }
    }
  }
  provideDefinition(
    document: vscode.TextDocument,
    position: vscode.Position,
  ): vscode.ProviderResult<vscode.Definition | vscode.DefinitionLink[]> {
    for (const handler of fileHandlers) {
      if (isMatch(document.uri.fsPath, handler.pattern)) {
        if (handler.process) {
          const ctx = createContext(document, position);
          const definitions = handler.process(ctx, this.rockide)?.definitions?.();
          return Promise.all(definitions ?? []);
        }
      }
    }
  }
  onDidChangeTextDocument({ document }: vscode.TextDocumentChangeEvent) {
    if (this.rockide.files.has(document.uri.fsPath)) {
      const text = document.getText();
      const json = JSONC.parse(text);
      this.rockide.files.set(document.uri.fsPath, json);
    }
  }
  async onDidCreateFiles({ files }: vscode.FileCreateEvent) {
    for (const uri of files) {
      if (!isMatch(uri.fsPath, `${baseGlob}/${projectGlob}/**/*.json`)) {
        continue;
      }
      if (uri.fsPath.endsWith(".json")) {
        await this.rockide.indexFile(uri);
        continue;
      }
      if (uri.fsPath.match(/\.(png|tga|fsb|ogg|wav)$/)) {
        this.rockide.indexAsset(uri);
        continue;
      }
    }
  }
  async onDidRenameFiles({ files }: vscode.FileRenameEvent) {
    for (const { oldUri, newUri } of files) {
      // If moved to outside project directory
      if (!isMatch(newUri.fsPath, `${baseGlob}/${projectGlob}/**/*.json`)) {
        this.rockide.files.delete(oldUri.fsPath);
        this.rockide.assets = this.rockide.assets.filter((v) => v.uri.fsPath === oldUri.fsPath);
        continue;
      }
      if (oldUri.fsPath.endsWith(".json")) {
        this.rockide.files.delete(oldUri.fsPath);
        this.rockide.jsonAssets = this.rockide.jsonAssets.filter((v) => v.uri.fsPath === oldUri.fsPath);
        await this.rockide.indexFile(newUri);
        continue;
      }
      this.rockide.assets = this.rockide.assets.filter((v) => v.uri.fsPath === oldUri.fsPath);
      if (newUri.fsPath.match(/\.(png|tga|fsb|ogg|wav)$/)) {
        this.rockide.indexAsset(newUri);
      }
    }
  }
  onDidDeleteFiles({ files }: vscode.FileDeleteEvent) {
    for (const uri of files) {
      if (isMatch(uri.fsPath, `${baseGlob}/${projectGlob}/**/*.json`)) {
        this.rockide.files.delete(uri.fsPath);
        this.rockide.assets = this.rockide.assets.filter((v) => v.uri.fsPath === uri.fsPath);
        this.rockide.jsonAssets = this.rockide.jsonAssets.filter((v) => v.uri.fsPath === uri.fsPath);
      }
    }
  }
}
