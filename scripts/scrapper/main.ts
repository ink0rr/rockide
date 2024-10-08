import fs from "node:fs/promises";
import path from "node:path/posix";

// config
const fileName = "block_identifier";

const snakeToCamel = (str: string) => str.replace(/_([a-z])/g, (g) => g[1].toUpperCase());
const exportName = snakeToCamel(fileName);

const filepath = path.join("node_modules", "bedrock-ts", "dist", "shared", "literals", `${fileName}.d.ts`);
const text = await fs.readFile(filepath, "utf-8");
const regex = /\|\s"(.*?)"/g;

let match;
const result: string[] = [];
while ((match = regex.exec(text))) {
  result.push(match[1]);
}

const dest = path.join("src", "literals", `${fileName}.ts`);
const data = `export const ${exportName} = [\n  "${result.join('",\n  "')}",\n];\n`;
await fs.mkdir(path.dirname(dest), { recursive: true });
await fs.writeFile(dest, data);
