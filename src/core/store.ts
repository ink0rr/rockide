import * as vscode from "vscode";
import { Reference } from "./reference";

export abstract class Store<T extends string = string> {
  abstract pattern: string;
  protected abstract store: Map<string, Reference[]>;
  abstract parse(uri: vscode.Uri): Promise<void> | void;
  get(key: T): Reference[] {
    return this.store.get(key) ?? [];
  }
  getFrom(uri: vscode.Uri, key: T) {
    return (this.store.get(key) ?? []).filter((v) => v.uri.fsPath === uri.fsPath);
  }
  delete(uri: vscode.Uri): void {
    for (const [id, values] of this.store) {
      const filtered = values.filter((v) => v.uri.fsPath !== uri.fsPath);
      this.store.set(id, filtered);
    }
  }
}