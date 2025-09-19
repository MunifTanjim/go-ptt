import { dirname, resolve as resolvePath } from "node:path";
import { fileURLToPath } from "url";

export const BIN_PATH = `${resolvePath(dirname(fileURLToPath(import.meta.url)), "../bin/ptt")}${process.platform === "win32" ? ".exe" : ""}`;
