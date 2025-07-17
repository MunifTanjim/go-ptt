import { ServiceError, credentials } from "@grpc/grpc-js";
import { ChildProcess, spawn } from "node:child_process";
import { existsSync } from "node:fs";
import { rm } from "node:fs/promises";
import { resolve as resolvePath } from "node:path";
import { setTimeout } from "node:timers/promises";
import {
  ParseResponse,
  ParseResponse_Result,
  ServiceClient,
} from "../proto/ptt";

export type ParseResult = ParseResponse_Result;

type PTTConfig = {
  socket: string;
};

const BIN_PATH = resolvePath(__dirname, "../bin/ptt");

export class PTTServer {
  socket: URL;

  #client!: ServiceClient;
  #process!: ChildProcess;

  constructor(conf: PTTConfig) {
    this.socket = URL.parse(
      conf.socket.startsWith("unix://")
        ? conf.socket
        : `unix://${resolvePath(conf.socket)}`,
    )!;
  }

  async start() {
    await rm(this.socket.pathname, { force: true });

    this.#process = spawn(
      BIN_PATH,
      ["server", "--socket", this.socket.pathname],
      { detached: true, stdio: "inherit" },
    );

    let timeLeft = 5000;
    let socketCreated = existsSync(this.socket.pathname);
    while (!socketCreated) {
      await setTimeout(200);
      timeLeft -= 200;
      if (timeLeft <= 0) {
        throw new Error(`failed start server`);
      }
      socketCreated = existsSync(this.socket.pathname);
    }

    process.on("exit", () => {
      this.stop();
    });

    this.#client = new ServiceClient(
      this.socket.toString(),
      credentials.createInsecure(),
    );
  }

  async stop() {
    this.#client.close();
    this.#process.kill();
  }

  async parse({
    torrent_titles,
  }: {
    torrent_titles: string[];
  }): Promise<ParseResult[]> {
    return new Promise<ParseResult[]>((resolve, reject) => {
      this.#client.parse(
        {
          torrent_titles: torrent_titles,
          normalize: true,
        },
        (err: ServiceError | null, response: ParseResponse) => {
          return err ? reject(err) : resolve(response.results);
        },
      );
    });
  }
}
