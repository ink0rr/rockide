import * as vscode from "vscode";
import { legend, SemanticToken } from "../semantics";

const molangSemantics: SemanticToken[] = [
  {
    pattern: /\b(q|v|t|c|query|variable|temp|context|math|array|geometry|material|texture)(?=\.)/gi,
    type: "class",
  },
  {
    pattern: /(?<=\b(q|v|t|c|query|variable|temp|context|math|array|geometry|material|texture))\.[\w.]+\b/gi,
    type: "function",
  },
  {
    pattern: /\bthis\b/gi,
    type: "macro",
  },
  {
    pattern: /(?<!\.)\b\d+(\.\d+)?(?!\.)f?\b/g,
    type: "number",
  },
  {
    pattern: /[=<>!&|,;?]+/g,
    type: "operator",
  },
  {
    pattern: /[\(\)\{\}\[\]]/g,
    type: "enumMember",
  },
  {
    pattern: /'[^']*'/g,
    type: "regexp",
  },
];

export class MolangProvider implements vscode.DocumentSemanticTokensProvider {
  provideDocumentSemanticTokens(document: vscode.TextDocument): vscode.ProviderResult<vscode.SemanticTokens> {
    const text = document.getText();
    const tokens = new vscode.SemanticTokensBuilder(legend);
    for (const { pattern, type, modifiers } of molangSemantics) {
      let match;
      while ((match = pattern.exec(text))) {
        const start = match.index;
        const length = match[0].length;
        const position = document.positionAt(start);
        const range = new vscode.Range(position, position.translate(0, length));
        tokens.push(range, type, modifiers);
      }
    }
    return tokens.build();
  }
}
