const {
  chmod,
  createWriteStream,
  existsSync,
  mkdirSync,
  unlinkSync,
} = require("node:fs");
const { join, relative, resolve } = require("node:path");
const { pipeline } = require("node:stream/promises");
const { extract } = require("tar");
const { promisify } = require("node:util");

const BINARY_NAME = "ptt";
const VERSION = "0.10.0";
const GITHUB_REPO = "MunifTanjim/go-ptt";

function getPlatformInfo() {
  const platform = process.platform;
  const arch = process.arch;

  /** @type {"darwin"|"linux"} */
  let platformName;
  /** @type {"amd64"|"arm64"} */
  let archName;

  switch (platform) {
    case "darwin":
      platformName = "darwin";
      break;
    case "linux":
      platformName = "linux";
      break;
    default:
      throw new Error(`Unsupported platform: ${platform}`);
  }

  switch (arch) {
    case "x64":
      archName = "amd64";
      break;
    case "arm64":
      archName = "arm64";
      break;
    default:
      throw new Error(`Unsupported architecture: ${arch}`);
  }

  return { platform: platformName, arch: archName };
}

async function downloadFile(url, dest) {
  const { default: fetch } = await import("node-fetch");

  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(
      `Failed to download: ${response.status} ${response.statusText}`,
    );
  }

  const fileStream = createWriteStream(dest);
  await pipeline(response.body, fileStream);
}

function makeExecutable(filePath) {
  return promisify(chmod)(filePath, 0o755);
}

async function install() {
  try {
    console.log("[ptt] Installing...");

    const { platform, arch } = getPlatformInfo();
    const archiveName = `${BINARY_NAME}_${VERSION}_${platform}_${arch}.tar.gz`;
    const downloadUrl = `https://github.com/${GITHUB_REPO}/releases/download/v${VERSION}/${archiveName}`;

    console.log(`[ptt] Downloading ${archiveName} for ${platform}/${arch}...`);

    const binDir = join(__dirname, "bin");
    if (!existsSync(binDir)) {
      mkdirSync(binDir, { recursive: true });
    }

    const archivePath = join(binDir, archiveName);
    await downloadFile(downloadUrl, archivePath);

    console.log("[ptt] Extracting Archive...");
    await extract({ file: archivePath, cwd: binDir });

    const binaryPath = join(binDir, BINARY_NAME);
    if (existsSync(binaryPath)) {
      await makeExecutable(binaryPath);
      console.log(
        `[ptt] Successfully Installed: ${relative(resolve(), binaryPath)}`,
      );
    } else {
      throw new Error(
        `[ptt] Binary ${BINARY_NAME} not found in extracted archive`,
      );
    }

    unlinkSync(archivePath);
  } catch (error) {
    console.error("[ptt] Installation Failed", error);
    process.exit(1);
  }
}

install();
