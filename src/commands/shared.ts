import { CompletionItem, CompletionItemKind, MarkdownString, SignatureHelp } from "vscode";
import { commands } from ".";
import { RockideContext } from "../context";
import { ParamInfo, ParamType } from "./types";

/**
 * TODO:
 * - Handle regex & parse location (~~~, ^^^, etc)
 * - idk why multiline doesn't work
 * - create separate context for non json files
 * - Merge signature location params
 */

/**
 * Convert ParamInfo to value
 * @param info
 * @returns
 */
function getParamValue(info: ParamInfo) {
  switch (info.type) {
    case ParamType.playerSelector:
      return ["@a", "@s", "@p", "@r"];
    case ParamType.entitySelector:
      return ["@a", "@s", "@p", "@r", "@e"];
    case ParamType.selectorWildcard:
      return ["@a", "@e", "@s", "@p", "@r", "*"];
    case ParamType.string:
      return '""';
    case ParamType.number:
      return "0";
    case ParamType.float:
      return "0.0";
    case ParamType.location:
      return ["~", "^", "0"];
    default:
      return info.value;
  }
}

/**
 * Convert ParamInfo to regex
 * @param info
 * @returns
 */
function getParamRegex(info: ParamInfo): RegExp {
  switch (info.type) {
    case ParamType.playerSelector:
      return /(@a|@s|@p|@r)(\[(.*?)\])?/g;
    case ParamType.entitySelector:
      return /(@a|@e|@s|@p|@r)(\[(.*?)\])?/g;
    case ParamType.selectorWildcard:
      return /(((@a|@e|@s|@p|@r)(\[(.*?)\])?)|\*)/g;
    case ParamType.string:
      return /("[^"]*"|\w+)/g;
    case ParamType.number:
      return /\d+/g;
    case ParamType.location:
      return /((~|\^|\d+)\.?)/g;
    default: {
      if (Array.isArray(info.value)) {
        return new RegExp(info.value.join("|"), "g");
      }
      return new RegExp(`${info.value}\\b`, "g");
    }
  }
}

/**
 * Convert ParamInfo to CompletionItem
 * @param info
 * @returns
 */
function getParamCompletion(info: ParamInfo): CompletionItem | CompletionItem[] {
  const getKind = (type: ParamType) => {
    switch (type) {
      case ParamType.keyword:
        return CompletionItemKind.Keyword;
      case ParamType.entitySelector:
      case ParamType.playerSelector:
      case ParamType.selectorWildcard:
        return CompletionItemKind.TypeParameter;
      case ParamType.tag:
        return CompletionItemKind.EnumMember;
      case ParamType.string:
        return CompletionItemKind.Value;
      case ParamType.number:
      case ParamType.float:
      case ParamType.yaw:
      case ParamType.pitch:
      case ParamType.vector3:
      case ParamType.location:
        return CompletionItemKind.Value;
      default:
        return CompletionItemKind.Keyword;
    }
  };
  const getDocs = (type: ParamType, value: string) => {
    switch (type) {
      case ParamType.playerSelector:
      case ParamType.selectorWildcard:
        {
          switch (value) {
            case "@a":
              return "All players.";
            case "@e":
              return "All entities.";
            case "@s":
              return "The entity running the command.";
            case "@p":
              return "The nearest player.";
            case "@r":
              return "A random player.";
            case "*":
              return "All players/entities.";
          }
        }
        break;
      default:
        return;
    }
  };
  const label = getParamValue(info);
  const createCompletionItem = (value: string) => {
    const documentation = getDocs(info.type, value) ?? info.documentation;
    const completion = new CompletionItem(value, getKind(info.type));
    completion.documentation = documentation;
    return completion;
  };

  return Array.isArray(label) ? label.map(createCompletionItem) : createCompletionItem(label);
}

export function commandCompletion(ctx: RockideContext): CompletionItem[] {
  const { document, position } = ctx;
  const line = document.lineAt(position.line).text;
  for (const { command, overloads } of commands) {
    if (line.startsWith("#")) {
      return [];
    }
    const match = line.match(new RegExp(`(\\b|\\/)${command}\\b`));
    if (match) {
      if (!overloads) {
        return [];
      }
      const [, ...args] = line
        .slice(match[0].length)
        .split(/\s+/g)
        .map((arg) => {
          const ok = new RegExp(/((~|\^)-?\d+())/g).test(arg);
          if (!ok) {
            return arg;
          }
          return arg.match(/((~|\^)-?\d+(\.\d+)?)/g) || [];
        })
        .flat();
      console.log(args);
      if (args.length === 1) {
        return overloads
          .map((overload) => getParamCompletion(overload.params[0]))
          .flat()
          .filter((v, i, s) => s.findIndex((obj) => JSON.stringify(obj) === JSON.stringify(v)) === i);
      }
      let tempOverloads = [...overloads];
      for (let i = 0; i < args.length; i++) {
        const arg = args[i];
        tempOverloads = tempOverloads.filter((overload) => {
          try {
            if (!overload.params[i]) {
              return false;
            }
            const regex = getParamRegex(overload.params[i]);
            if (i === args.length - 1) {
              return !regex.test(arg);
            }
            return regex.test(arg);
          } catch (e) {
            console.error(e);
          }
        });
      }
      const completions = tempOverloads
        .map((overload) => getParamCompletion(overload.params[args.length - 1]))
        .flat()
        .filter((v, i, s) => s.findIndex((obj) => JSON.stringify(obj) === JSON.stringify(v)) === i);
      return completions;
    }
  }
  return commands.map(({ command, documentation }) => {
    const completion = new CompletionItem(command, CompletionItemKind.Class);
    const markdown = new MarkdownString(`## ${command}\n\n${documentation}`);
    completion.documentation = markdown;
    return completion;
  });
}

export function signatureHelper(ctx: RockideContext): SignatureHelp | undefined {
  const { document, position } = ctx;
  const line = document.lineAt(position.line).text;
  const signature: SignatureHelp = new SignatureHelp();

  for (const { command, overloads, documentation } of commands) {
    if (line.startsWith("#")) {
      return;
    }
    const match = line.match(new RegExp(`^${command}\\b`));
    if (match) {
      if (!overloads) {
        signature.activeSignature = 0;
        signature.activeParameter = 0;
        signature.signatures = [
          {
            label: command,
            documentation,
            parameters: [],
          },
        ];
        return signature;
      }
      const [, ...args] = line
        .slice(match[0].length)
        .split(/\s+/g)
        .map((arg) => {
          const ok = new RegExp(/((~|\^)-?\d+())/g).test(arg);
          if (!ok) {
            return arg;
          }
          return arg.match(/((~|\^)-?\d+(\.\d+)?)/g) || [];
        })
        .flat();
      let tempOverloads = [...overloads];
      for (let i = 0; i < args.length; i++) {
        const arg = args[i];
        tempOverloads = tempOverloads.filter((overload) => {
          try {
            if (!overload.params[i]) {
              return false;
            }
            const regex = getParamRegex(overload.params[i]);
            if (i === args.length - 1) {
              return !regex.test(arg);
            }
            return regex.test(arg);
          } catch (e) {
            console.error(e);
          }
        });
      }
      const sigs = tempOverloads.map((overload) => {
        const params = overload.params.map((param) => {
          let label = param.signatureValue ?? param.value;
          if (Array.isArray(label)) {
            label = label.join("|");
          }
          return {
            label,
            documentation: param.documentation,
          };
        });

        return {
          label: `${command} ${params.map((param) => param.label).join(" ")}`,
          documentation,
          parameters: params.map(({ documentation, label }) => {
            return {
              label,
              documentation,
            };
          }),
        };
      });
      // const sigs: SignatureInformation[] = [];
      // tempOverloads.forEach((overload) => {
      //   const params = overload.params.map((param) => {
      //     let label = param.value;
      //     if (Array.isArray(label)) {
      //       label = label.join("|");
      //     }
      //     return {
      //       label,
      //       documentation: param.documentation,
      //     };
      //   });

      //   sigs.push({
      //     label: `${command} ${params.map((param) => param.label).join(" ")}`,
      //     documentation,
      //     parameters: params.map((param) => {
      //       let label = param.label;
      //       if (Array.isArray(label)) {
      //         label = label.join("|");
      //       }
      //       return {
      //         label,
      //         documentation: param.documentation,
      //       };
      //     }),
      //   });
      // });
      signature.activeSignature = 0;
      signature.activeParameter = args.length - 1;
      signature.signatures = sigs;
      return signature;
    }
  }
  return;
}

export namespace Parameter {
  export const position: ParamInfo[] = [
    {
      value: "<x>",
      type: ParamType.location,
    },
    {
      value: "<y>",
      type: ParamType.location,
    },
    {
      value: "<z>",
      type: ParamType.location,
    },
  ];
  export const rotXY: ParamInfo[] = [
    {
      value: "<pitch>",
      type: ParamType.pitch,
    },
    {
      value: "<yaw>",
      type: ParamType.yaw,
    },
  ];
  export const rotYX: ParamInfo[] = [
    {
      value: "<yaw>",
      type: ParamType.yaw,
    },
    {
      value: "<pitch>",
      type: ParamType.pitch,
    },
  ];
}
