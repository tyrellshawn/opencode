# Local Binary Build Setup

## Summary

Successfully configured the opencode project to build and use a local compiled binary instead of running TypeScript through `bun run`.

## What Was Done

### 1. **Created build-local.ts script** (`packages/opencode/script/build-local.ts`)

- Builds opencode for the current platform/architecture only (not all platforms like the main build script)
- Uses Bun's `--compile` target to create a native binary
- Stores the binary in `.bin/opencode` (or `.bin/opencode.exe` on Windows)
- Uses `@opentui/solid/scripts/solid-plugin` for JSX compilation
- Properly resolves dependencies from the workspace root node_modules

### 2. **Key Implementation Details**

**Path Resolution Issue & Solution:**

- Initial problem: The script is at `packages/opencode/script/` but dependencies are installed at the workspace root
- Tried relative imports from `../node_modules` and `../../node_modules` - both failed
- Solution: Use absolute path resolution:
  ```typescript
  const rootDir = path.resolve(__dirname, "../../..") // Go up 3 levels
  const solidPluginPath = path.resolve(rootDir, "node_modules/@opentui/solid/scripts/solid-plugin.ts")
  const solidPlugin = await import(solidPluginPath)
  ```

**Version Handling:**

- Removed dependency on `@opencode-ai/script` which had strict bun version checking
- Implemented simpler version logic directly in the build script:
  ```typescript
  const CHANNEL = process.env["OPENCODE_CHANNEL"] ?? (await $`git branch --show-current`.text().then((x) => x.trim()))
  const VERSION = process.env["OPENCODE_VERSION"] ?? `0.0.0-${CHANNEL}`
  ```

**Build Configuration:**

- Used Bun's native compilation with `--compile` flag
- Includes:
  - SolidJS plugin for JSX transformation
  - Tree-sitter parser worker
  - UI worker
  - Source maps
  - Proper environment defines (OPENCODE_VERSION, OPENCODE_CHANNEL, etc.)

### 3. **Updated Fish Config** (`~/.config/fish/config.fish`)

- Created `opencode` function that calls the binary directly
- Includes fallback error message if binary doesn't exist
- Set `OPENCODE_BIN_PATH` environment variable for debugging

### 4. **Added npm script** (`packages/opencode/package.json`)

- Added `"build:local": "./script/build-local.ts"` script for easy building
- Can now run `bun build:local` from the opencode package directory

## Build Results

- Binary size: ~128 MB (uncompressed)
- Location: `/home/tyrellshawn/agents/opencode/packages/opencode/.bin/opencode`
- Version: `0.0.0-dev-nov-2025`
- Execution: Direct binary call (no bun runtime needed)

## Usage

### Build the binary:

```bash
cd /home/tyrellshawn/agents/opencode/packages/opencode
bun build:local
```

### Use the function:

```bash
# Any new shell will have the opencode function available
opencode --version
opencode .
```

## Gotchas & Lessons

1. **Workspace structure**: Dependencies are installed at root level, not per-package. Bun's monorepo setup requires going 3 levels up from the script directory.

2. **Version checking**: The `@opencode-ai/script` module enforces strict version checking. For local development builds, it's better to avoid this dependency or make it optional.

3. **SolidJS plugin**: The plugin needs to be loaded dynamically via `await import()` with absolute paths in monorepo setups, not ES6 static imports.

4. **Bun versions**: The project specifies `bun@1.3.2` but bun 1.3.3 is available. Both work fine for compilation despite the version check in Script module.

5. **Tree-sitter dependencies**: The build properly includes the parser worker from `@opentui/core` which contains platform-specific binaries.

## Related Files

- `/home/tyrellshawn/.config/fish/config.fish` - Fish shell function
- `/home/tyrellshawn/agents/opencode/packages/opencode/script/build-local.ts` - Build script
- `/home/tyrellshawn/agents/opencode/packages/opencode/package.json` - npm script definition

## What Works

✅ Binary builds successfully  
✅ Binary executes correctly  
✅ Fish config function calls it  
✅ --version flag works  
✅ All dependencies properly included (SolidJS, tree-sitter, etc.)

## What Could Be Improved

- Binary size could be reduced with stripping
- Build time is ~1-2 minutes (acceptable for local dev)
- Could add auto-rebuild on source changes
- Could compress binary after building
