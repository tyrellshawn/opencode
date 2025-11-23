#!/usr/bin/env bun

import path from "path"
import fs from "fs"
import { $ } from "bun"
import { fileURLToPath } from "url"

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
const scriptDir = path.resolve(__dirname)
const pkgDir = path.resolve(__dirname, "..")
const rootDir = path.resolve(__dirname, "../../..")

process.chdir(pkgDir)

// Import from root node_modules using absolute path
const solidPluginPath = path.resolve(rootDir, "node_modules/@opentui/solid/scripts/solid-plugin.ts")
const solidPlugin = await import(solidPluginPath)

const pkg = await Bun.file(path.resolve(pkgDir, "package.json")).json()

// Get version info without version check
const CHANNEL = process.env["OPENCODE_CHANNEL"] ?? (await $`git branch --show-current`.text().then((x) => x.trim()))
const VERSION = process.env["OPENCODE_VERSION"] ?? `0.0.0-${CHANNEL}`

const platformMap: Record<string, string> = {
    darwin: "darwin",
    linux: "linux",
    win32: "windows",
}
const archMap: Record<string, string> = {
    x64: "x64",
    arm64: "arm64",
    arm: "arm",
}

const platform = platformMap[process.platform] || process.platform
const arch = archMap[process.arch] || process.arch

console.log(`Building opencode binary for ${platform}-${arch}...`)

await $`mkdir -p .bin`

const parserWorker = fs.realpathSync(path.resolve(rootDir, "node_modules/@opentui/core/parser.worker.js"))
const workerPath = "./src/cli/cmd/tui/worker.ts"
const binary = platform === "windows" ? "opencode.exe" : "opencode"
const outfile = `.bin/${binary}`

await Bun.build({
    conditions: ["browser"],
    tsconfig: "./tsconfig.json",
    plugins: [solidPlugin.default],
    sourcemap: "external",
    compile: {
        target: `bun-${platform}-${arch}` as any,
        outfile,
        execArgv: [`--user-agent=opencode/${VERSION}`, `--env-file=""`, `--`],
        windows: {},
    },
    entrypoints: ["./src/index.ts", parserWorker, workerPath],
    define: {
        OPENCODE_VERSION: `'${VERSION}'`,
        OTUI_TREE_SITTER_WORKER_PATH: "/$bunfs/root/" + path.relative(pkgDir, parserWorker).replaceAll("\\", "/"),
        OPENCODE_WORKER_PATH: workerPath,
        OPENCODE_CHANNEL: `'${CHANNEL}'`,
    },
})

await $`chmod +x ${outfile}`

console.log(`✓ Built binary at ${outfile}`)
