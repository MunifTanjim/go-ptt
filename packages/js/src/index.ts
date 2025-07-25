import { ServiceError, credentials } from "@grpc/grpc-js";
import { ChildProcess, spawn } from "node:child_process";
import { rm } from "node:fs/promises";
import { resolve as resolvePath } from "node:path";
import { setTimeout } from "node:timers/promises";
import {
  ParseRequest,
  ParseResponse,
  ParseResponse_Result,
  PingRequest,
  PingResponse,
  ServiceClient,
} from "../proto/ptt";

export type ParseResult = ParseResponse_Result;

type PTTConfig = {
  /**
   * For Windows, `unix` domain socket is not supported.
   */
  network: "tcp" | "unix";
  /**
   * For `unix` network, should be path for domain socket.
   * For `tcp` network, should be `host:port` or `:port`.
   */
  address: string;
};

const BIN_PATH = `${resolvePath(__dirname, "../bin/ptt")}${process.platform === "win32" ? ".exe" : ""}`;

export class PTTServer {
  network: PTTConfig["network"];
  address: PTTConfig["address"];

  #client!: ServiceClient;
  #process!: ChildProcess;

  constructor(conf: PTTConfig) {
    this.network = conf.network;
    this.address = conf.address;
    switch (this.network) {
      case "tcp":
        if (this.address.startsWith(":")) {
          this.address = `localhost${this.address}`;
        }
        break;
      case "unix":
        if (!this.address.startsWith("/") && !this.address.startsWith("./")) {
          this.address = resolvePath(`./${this.address}`);
        }
        break;
    }
  }

  async ping({ message }: PingRequest): Promise<PingResponse> {
    return new Promise<PingResponse>((resolve, reject) => {
      this.#client.ping(
        { message },
        (err: ServiceError | null, response: PingResponse) => {
          return err ? reject(err) : resolve(response);
        },
      );
    });
  }

  async start() {
    if (this.network === "unix") {
      await rm(this.address, { force: true });
    }

    this.#process = spawn(
      BIN_PATH,
      ["server", "--network", this.network, "--address", this.address],
      { detached: true, stdio: "inherit" },
    );

    let timeLeft = 5000;
    let isReady = false;
    while (!isReady) {
      try {
        if (!this.#client) {
          this.#client = new ServiceClient(
            this.network === "tcp"
              ? `${this.address}`
              : `${this.network}://${this.address}`,
            credentials.createInsecure(),
          );
        }

        await this.ping({ message: "" });
        isReady = true;
      } catch (err) {
        await setTimeout(200);
        timeLeft -= 200;
        if (timeLeft <= 0) {
          throw new Error(`failed start server`, { cause: err });
        }
      }
    }

    process.on("SIGINT", () => {
      this.stop();
    });
  }

  async stop() {
    this.#client.close();
    this.#process.kill();
  }

  async parse({
    torrent_titles,
    normalize = false,
  }: {
    torrent_titles: ParseRequest["torrent_titles"];
    normalize?: ParseRequest["normalize"];
  }): Promise<ParseResult[]> {
    return new Promise<ParseResult[]>((resolve, reject) => {
      this.#client.parse(
        {
          torrent_titles,
          normalize,
        },
        (err: ServiceError | null, response: ParseResponse) => {
          return err ? reject(err) : resolve(response.results);
        },
      );
    });
  }
}
