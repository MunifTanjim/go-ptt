#!/usr/bin/env node

import { resolve as resolvePath } from "node:path";
import { spawn } from "child_process";

const BIN_PATH = resolvePath(__dirname, "../bin/ptt");

const child = spawn(BIN_PATH, process.argv.slice(2), {
  stdio: "inherit",
});

child.on("close", (code) => {
  process.exit(code);
});

child.on("error", (error) => {
  console.error(error);
  process.exit(1);
});
