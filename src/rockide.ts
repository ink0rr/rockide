import * as JSONC from "jsonc-parser";
import { isMatch } from "micromatch";
import { relative } from "path";
import * as vscode from "vscode";
import { bpGlob, NullNode, projectGlob, rpGlob } from "./constants";
import { fileHandlers } from "./handlers";

export type IndexedData = {
  path: string;
  root: JSONC.Node;
  values: string[];
};

export type AssetData = {
  uri: vscode.Uri;
  bedrockPath: string;
};

export class Rockide {
  diagnostics = vscode.languages.createDiagnosticCollection("rockide");
  files = new Map<string, JSONC.Node>();
  assets: AssetData[] = [];
  jsonAssets: AssetData[] = [];
  mcfunctions = new Map<string, string>(); // path, content

  async checkWorkspace() {
    for (const path of await vscode.workspace.findFiles("**/manifest.json")) {
      const file = await vscode.workspace.openTextDocument(path);
      const json = JSONC.parse(file.getText());
      if ("header" in json && "format_version" in json) {
        continue;
      }
      return false;
    }
    return true;
  }

  async indexWorkspace() {
    if (!vscode.workspace.workspaceFolders) {
      return;
    }
    if (vscode.workspace.workspaceFolders.length > 1) {
      return vscode.window.showInformationMessage("Multiple workspace is currently not supported.");
    }
    const workspace = vscode.workspace.workspaceFolders[0];
    this.files.clear();
    this.assets = [];
    vscode.window.withProgress(
      { title: "Indexing", location: vscode.ProgressLocation.Notification },
      async (progress) => {
        const fileList = await vscode.workspace.findFiles(`**/${projectGlob}/**/*.json`, "{.*,build}/**");
        const increment = 100 / fileList.length;
        for (const uri of fileList) {
          progress.report({ message: relative(workspace.uri.fsPath, uri.fsPath), increment });
          await this.indexFile(uri);
        }
        const mcfunctionList = await vscode.workspace.findFiles(`**/${bpGlob}/functions/*.mcfunction`, "{.*,build}/**");
        for (const uri of mcfunctionList) {
          this.indexMcfunction(uri);
        }
        const assetList = await vscode.workspace.findFiles(`**/${rpGlob}/**/*.{png,tga,fsb,ogg,wav}`, "{.*,build}/**");
        for (const uri of assetList) {
          this.indexAsset(uri);
        }
      },
    );
  }

  async indexFile(uri: vscode.Uri) {
    for (const handler of fileHandlers) {
      if (!handler.index || !isMatch(uri.fsPath, handler.pattern)) {
        continue;
      }
      if (handler.index === "parse") {
        const document = await vscode.workspace.openTextDocument(uri);
        const root = JSONC.parseTree(document.getText()) ?? NullNode;
        this.files.set(uri.fsPath, root);
        break;
      }
      const path = uri.fsPath.replaceAll("\\", "/").split(/(behavior_pack|[^\\/]*?bp|bp_[^\\/]*?)\//i)[2];
      if (path) {
        this.jsonAssets.push({
          uri,
          bedrockPath: path,
        });
      }
      break;
    }
  }

  async indexMcfunction(uri: vscode.Uri) {
    const document = await vscode.workspace.openTextDocument(uri);
    this.mcfunctions.set(uri.fsPath, document.getText());
  }

  indexAsset(uri: vscode.Uri) {
    const path = uri.fsPath.replaceAll("\\", "/").split(/(resource_pack|[^\\/]*?rp|rp_[^\\/]*?)\//i)[2];
    if (path) {
      this.assets.push({
        uri,
        bedrockPath: path.replace(/\.\w+$/, ""),
      });
    }
  }

  getAnimations(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${bpGlob}/{animations,animation_controllers}/**/*.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        return { path, root, values: Object.keys(json.animations ?? json.animation_controllers ?? {}) };
      });
  }

  getClientAnimations(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${rpGlob}/{animations,animation_controllers}/**/*.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        return { path, root, values: Object.keys(json.animations ?? json.animation_controllers ?? {}) };
      });
  }

  getGeometries(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${rpGlob}/models/**/*.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        if (Array.isArray(json["minecraft:geometry"])) {
          return {
            path,
            root,
            values: json["minecraft:geometry"]
              .map((geo) => geo.description?.identifier)
              .filter((key: string | undefined) => key?.startsWith("geometry.")),
          };
        }
        delete json.format_version;
        return {
          path,
          root,
          values: Object.keys(json).filter((key) => key.startsWith("geometry.")),
        };
      });
  }

  getRenderControllers(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${rpGlob}/render_controllers/**/*.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        return { path, root, values: Object.keys(json.render_controllers ?? {}) };
      });
  }

  getParticles(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${rpGlob}/particles/**/*.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        const identifier = json.particle_effect?.description?.identifier;
        return { path, root, values: identifier ? [identifier] : [] };
      });
  }

  getItemIcons(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${rpGlob}/textures/item_texture.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        return { path, root, values: Object.keys(json.texture_data) };
      });
  }

  getSoundDefinitions(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${rpGlob}/sounds/sound_definitions.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        return { path, root, values: Object.keys(json.sound_definitions) };
      });
  }

  getManifests(): IndexedData[] {
    return [...this.files]
      .filter(([path]) => isMatch(path, `**/${projectGlob}/manifest.json`))
      .map(([path, root]) => {
        const json = JSONC.getNodeValue(root);
        const uuid = json?.header?.uuid;
        return { path, root, values: uuid ? [uuid] : uuid };
      });
  }

  getTextures() {
    return this.assets.filter(({ bedrockPath: path }) => path.startsWith("textures/"));
  }

  getSounds() {
    return this.assets.filter(({ bedrockPath: path }) => path.startsWith("sounds/"));
  }

  getLootTables() {
    return this.jsonAssets.filter(({ bedrockPath: path }) => path.startsWith("loot_tables/"));
  }

  getTradeTables() {
    return this.jsonAssets.filter(({ bedrockPath: path }) => path.startsWith("trading/"));
  }
}
