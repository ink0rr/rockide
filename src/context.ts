import * as JSONC from "jsonc-parser";
import * as vscode from "vscode";
import { NullNode } from "./constants";
import { RockideProcess } from "./handlers/types";

export function createContext(document: vscode.TextDocument, position: vscode.Position) {
  const text = document.getText();
  const offset = document.offsetAt(position);
  const location = JSONC.getLocation(text, offset);
  const node = location.previousNode ?? NullNode;
  const path = location.path;
  const nodeValue = JSONC.getNodeValue(node) ?? node.value ?? "";
  // if (typeof nodeValue !== "string") {
  // throw new Error("nodeValue is not a string");
  // }
  return {
    document,
    position,
    path,
    node,
    nodeValue,
    rootJson() {
      return JSONC.parse(text);
    },
    findNode(path: JSONC.JSONPath) {
      const root = JSONC.parseTree(text);
      if (!root) {
        return;
      }
      return JSONC.findNodeAtLocation(root, path);
    },
    getKeys(path: JSONC.JSONPath) {
      let json = this.rootJson();
      for (const key of path) {
        if (key in json) {
          json = json[key];
        } else {
          return [];
        }
      }
      if (typeof json === "object") {
        return Object.keys(json);
      }
      return [];
    },
    /**
     * ```json
     * {
     *   "key": "value"
     * }
     * ```
     */
    matchField(key: string) {
      return node.type.match(/(string|null)/) && path.at(-1) === key && !location.isAtPropertyKey;
    },
    /**
     * ```json
     * {
     *   "root": {
     *     "key": "value"
     *   }
     * }
     * ```
     */
    matchProperty(root: string, key?: string) {
      const parentKey = path.at(-1);
      return (
        node.type.match(/(string|null)/) &&
        path.at(-2) === root &&
        (key ? parentKey === key : parentKey !== "number") &&
        !location.isAtPropertyKey
      );
    },
    /**
     * ```json
     * {
     *   "key": [
     *     "value"
     *    ]
     * }
     * ```
     */
    matchArray(key: string, hasConditional?: boolean) {
      return (
        (hasConditional && this.matchConditionalArray(key)) ||
        (node.type.match(/(string|null)/) && typeof path.at(-1) === "number" && path.at(-2) === key)
      );
    },
    /**
     * ```json
     * {
     *   "key": [
     *     { "value": "..." },
     *    ]
     * }
     * ```
     */
    matchConditionalArray(key: string) {
      return location.isAtPropertyKey && typeof path.at(-2) === "number" && path.at(-3) === key;
    },
    /**
     * ```json
     * {
     *   "root": [
     *     { "key": "value" },
     *    ]
     * }
     * ```
     */
    matchArrayObject(root: string, key?: string) {
      return (key ? this.matchField(key) : true) && typeof path.at(-2) === "number" && path.at(-3) === root;
    },

    createCompletion(value: string | vscode.CompletionItem) {
      if (value instanceof vscode.CompletionItem) {
        return value;
      }
      const completion = new vscode.CompletionItem(value, vscode.CompletionItemKind.Variable);
      completion.range = document.getWordRangeAtPosition(position, /[\w.:/]+/);
      if (node.type === "null") {
        completion.insertText = `"${value}"`;
        if (location.isAtPropertyKey) {
          completion.insertText += `: ""`;
          completion.command = {
            title: "Move cursor",
            command: "cursorLeft",
          };
        }
      }
      return completion;
    },
    async createDefinition(path: string, targetNode?: JSONC.Node): Promise<vscode.LocationLink> {
      const originSelectionRange = new vscode.Range(
        document.positionAt(node.offset),
        document.positionAt(node.offset + node.length),
      );
      if (!path.endsWith(".json")) {
        const position = new vscode.Position(0, 0);
        return {
          originSelectionRange,
          targetRange: new vscode.Range(position, position),
          targetUri: vscode.Uri.file(path),
        };
      }
      const targetDocument = await vscode.workspace.openTextDocument(path);
      if (!targetNode) {
        const position = new vscode.Position(0, 0);
        return {
          originSelectionRange,
          targetRange: new vscode.Range(position, position),
          targetUri: targetDocument.uri,
        };
      }
      if (targetNode.parent) {
        targetNode = targetNode.parent;
      }
      return {
        originSelectionRange,
        targetRange: new vscode.Range(
          targetDocument.positionAt(targetNode.offset),
          targetDocument.positionAt(targetNode.offset + targetNode.length),
        ),
        targetUri: targetDocument.uri,
      };
    },
    localRef(path: JSONC.JSONPath): RockideProcess {
      return {
        completions: () => this.getKeys(path),
        definitions: () => {
          let target = this.findNode(path.concat([this.nodeValue]));
          if (!target) {
            return;
          }
          if (target.parent) {
            target = target.parent;
          }
          return [
            {
              originSelectionRange: new vscode.Range(
                document.positionAt(node.offset),
                document.positionAt(node.offset + node.length),
              ),
              targetRange: new vscode.Range(
                document.positionAt(target.offset),
                document.positionAt(target.offset + target.length),
              ),
              targetUri: document.uri,
            },
          ];
        },
      };
    },
  };
}

export type RockideContext = NonNullable<ReturnType<typeof createContext>>;
