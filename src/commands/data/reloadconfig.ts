import { CommandInfo } from "../types";

const reloadconfig: CommandInfo = {
  command: "reloadconfig",
  documentation: "Reloads configuration files relating to variables, secrets, permissions, etc.",
  overloads: [
    {
      params: [],
    },
  ],
};
export default reloadconfig;
