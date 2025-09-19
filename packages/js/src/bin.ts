#!/usr/bin/env node

import { spawn } from "child_process";
import { BIN_PATH } from "./const";

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
