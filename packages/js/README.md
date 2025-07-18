# go-ptt

Golang - Parse Torrent Title

## Installation

```sh
# using pnpm:
pnpm add go-ptt

# using npm:
npm install --save go-ptt

# using yarn:
yarn add go-ptt
```

## Usage

**Basic Usage:**

```js
import { PTTServer } from "go-ptt";

const server = new PTTServer({
  network: "unix", // or "tcp"
  address: "ptt.sock", // or ":8888"
});

await server.start();

const torrent_titles = ["Friends S01E01"];

const results = await server.parse({
  torrent_titles,
  normalize: true,
});

torrent_titles.forEach((torrent_title, idx) => {
  const result = results[idx];
  if (result.err) {
    console.error(result.err);
  } else {
    console.log(result);
  }
});

await server.stop();
```

## License

Licensed under the MIT License. Check the [LICENSE](./LICENSE) file for details.
