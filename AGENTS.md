# AGENTS.md
## Commands
- **Install**: `bun install`
- **Run**: `bun run index.ts`
- **Typecheck**: `bun run typecheck`
- **Test all**: `bun test`
- **Test single**: `bun test test/tool/tool.test.ts`
- **Generate API client**: `cd packages/tui && go generate ./pkg/client/`
## Formatting & Linting
- Prettier configured (semi: false)
- `npx prettier --check .` / `npx prettier --write .`
## Code Style
- Bun + TS ESM; relative imports; named imports
- camelCase vars/funcs; PascalCase types/classes
- Zod schemas for validation; avoid `any`; `const` > `let`
- Single responsibility; avoid `try`/`catch` & `else`
## Architecture
- Tools implement `Tool.Info` interface; use `App.provide()` for DI; `Log.create`, `Storage`
## Cursor & Copilot
- none detected

## Style

- prefer single word variable/function names
- avoid try catch where possible - prefer to let exceptions bubble up
- avoid else statements where possible
- do not make useless helper functions - inline functionality unless the
  function is reusable or composable
- prefer Bun apis

## Workflow

- you can regenerate the golang sdk by calling ./scripts/stainless.ts
- we use bun for everything
