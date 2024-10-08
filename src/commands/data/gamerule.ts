import { CommandInfo, ParamType } from "../types";

const gamerule: CommandInfo = {
  command: "gamerule",
  documentation: "Sets or queries a game rule value.",
  overloads: [
    {
      params: [],
    },
    {
      params: [
        {
          value: [
            "commandblockoutput",
            "dodaylightcycle",
            "doentitydrops",
            "dofiretick",
            "recipesunlock",
            "dolimitedcrafting",
            "domobloot",
            "domobspawning",
            "dotiledrops",
            "doweathercycle",
            "drowningdamage",
            "falldamage",
            "firedamage",
            "keepinventory",
            "mobgriefing",
            "pvp",
            "showcoordinates",
            "showdaysplayed",
            "naturalregeneration",
            "tntexplodes",
            "sendcommandfeedback",
            "doinsomnia",
            "commandblocksenabled",
            "doimmediaterespawn",
            "showdeathmessages",
            "showtags",
            "freezedamage",
            "respawnblocksexplode",
            "showbordereffect",
            "showrecipemessages",
            "projectilescanbreakblocks",
            "tntexplosiondropdecay",
          ],
          type: ParamType.keyword,
        },
        {
          value: ["true", "false"],
          type: ParamType.keyword,
        },
      ],
    },
    {
      params: [
        {
          value: [
            "maxcommandchainlength",
            "randomtickspeed",
            "functioncommandlimit",
            "spawnradius",
            "playerssleepingpercentage",
          ],
          type: ParamType.keyword,
        },
        {
          value: "int",
          type: ParamType.keyword,
        },
      ],
    },
  ],
};
export default gamerule;
