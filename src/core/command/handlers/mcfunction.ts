import { bpGlob } from "../../../constants";
import { commandCompletion, signatureHelper } from "../shared";
import { ParamType } from "../types";
import { CommandHandler } from "./_types";

export const mcfunctionHandler: CommandHandler = {
  pattern: `**/${bpGlob}/functions/**/*.mcfunction`,
  index: true,
  process(ctx, rockide) {
    return {
      completions() {
        // console.log(ctx.getCommandsV2());
        return commandCompletion(ctx, rockide);
      },
      signature() {
        return signatureHelper(ctx, rockide);
      },
      definitions() {
        const mcfunctions = Array.from(rockide.getMcfunctions());
        const commandSequences = ctx.getCommandsV2();
        if (!commandSequences.length) {
          return;
        }
        const { args } = commandSequences[commandSequences.length - 1];
        const currentWord = ctx.getCurrentWord();
        if (!currentWord) {
          return;
        }
        let { range, text } = currentWord;
        text = text.replace(/\"/g, "");
        const path = mcfunctions.find((path) => path.endsWith(`${text}.mcfunction`));
        if (!path) {
          return;
        }
        return args
          .filter((arg) => arg.type === ParamType.RockideMcfunction)
          .map(() => ctx.createDefinition(path, range));
      },
    };
  },
};
