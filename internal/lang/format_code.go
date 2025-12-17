package lang

type FormatCode struct {
	Description string // optional description
}

var FormatCodes = map[rune]FormatCode{
	'a': {
		Description: "Green",
	},
	'b': {
		Description: "Aqua",
	},
	'c': {
		Description: "Red",
	},
	'd': {
		Description: "Light Purple",
	},
	'e': {
		Description: "Yellow",
	},
	'f': {
		Description: "White",
	},
	'g': {
		Description: "Minecoin Gold",
	},
	'h': {
		Description: "Material Quartz",
	},
	'i': {
		Description: "Material Iron",
	},
	'j': {
		Description: "Material Netherite",
	},
	'k': {
		Description: "Obfuscated / MTS",
	},
	'l': {
		Description: "Bold",
	},
	'm': {
		Description: "Material Redstone",
	},
	'n': {
		Description: "Material Copper",
	},
	'o': {
		Description: "Italic",
	},
	'p': {
		Description: "Material Gold",
	},
	'q': {
		Description: "Material Emerald",
	},
	'r': {
		Description: "Reset",
	},
	's': {
		Description: "Material Diamond",
	},
	't': {
		Description: "Material Lapis Lazuli",
	},
	'u': {
		Description: "Material Amethyst",
	},
	'v': {
		Description: "Material Resin",
	},
	'0': {
		Description: "Black",
	},
	'1': {
		Description: "Dark Blue",
	},
	'2': {
		Description: "Dark Green",
	},
	'3': {
		Description: "Dark Aqua",
	},
	'4': {
		Description: "Dark Red",
	},
	'5': {
		Description: "Dark Purple",
	},
	'6': {
		Description: "Gold",
	},
	'7': {
		Description: "Gray",
	},
	'8': {
		Description: "Dark Gray",
	},
	'9': {
		Description: "Blue",
	},
}
