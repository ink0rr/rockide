import { commandCompletion, signatureHelper } from "../../commands/shared";
import { bpGlob } from "../../constants";
import { RockideHandler } from "../types";

export const mcfunctionHandler: RockideHandler = {
  pattern: `**/${bpGlob}/functions/**/*.mcfunction`,
  index: "mcfunction",
  process(ctx) {
    return {
      completions() {
        return commandCompletion(ctx);
      },
      signature() {
        return signatureHelper(ctx);
      },
    };
  },
};
