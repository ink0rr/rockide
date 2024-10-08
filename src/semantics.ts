import * as vscode from "vscode";

export type SemanticToken = {
  pattern: RegExp;
  type: string;
  modifiers?: string[];
};

export const molangSemantics: SemanticToken[] = [
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

export const legend = new vscode.SemanticTokensLegend([
  "class",
  "enumMember",
  "function",
  "macro",
  "number",
  "operator",
  "regexp",
  "string",
  "variable",
  "comment",
  "keyword",
]);
