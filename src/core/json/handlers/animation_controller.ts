import { bpGlob } from "../../../constants";
import { getMolangCompletions } from "../../molang/handlers";
import { JsonHandler } from "./_type";

export const animationControllerHandler: JsonHandler = {
  pattern: `**/${bpGlob}/animation_controllers/**/*.json`,
  index: "parse",
  process(ctx) {
    const id = ctx.path[1];
    if (ctx.matchConditionalArray("transitions")) {
      return ctx.localRef(["animation_controllers", id, "states"]);
    }
    if (
      ctx.matchArray("on_entry") ||
      ctx.matchArray("on_exit") ||
      ctx.matchArrayObject("animations") ||
      ctx.matchArrayObject("transitions")
    ) {
      return {
        completions: () => getMolangCompletions(ctx.document, ctx.position),
      };
    }
  },
};
