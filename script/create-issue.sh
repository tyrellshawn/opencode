#!/bin/bash
# Create GitHub issue for opencode setup fix

gh issue create \
  --title "Fix: Clean up multiple opencode installations and use local development build" \
  --body "## Problem
Multiple outdated opencode installations were conflicting with each other:
- \`/home/tyrellshawn/.bun/bin/opencode\` (old npm install)
- \`/home/tyrellshawn/.opencode/bin/opencode\` (old install script)
- \`/usr/local/bin/opencode\` (old compiled binary from June)

This caused \`which opencode\` to resolve to stale binaries instead of the local development version.

## Solution
1. Removed \`/home/tyrellshawn/.bun/bin/opencode\` (old npm install)
2. Removed \`/home/tyrellshawn/.opencode/bin/opencode\` (old install script)
3. Removed \`~/.bun/bin\` from PATH in fish config
4. Added code to remove \`~/.opencode/bin\` from \`fish_user_paths\`
5. Configured fish function to run local development build: \`bun run --conditions=development /home/tyrellshawn/agents/opencode/packages/opencode/src/index.ts\`

## Result
✅ \`opencode\` command now properly resolves to the fish function
✅ Runs local development build from source instead of stale binaries
✅ Ready for local development with \`opencode .\`

## Files Modified
- \`/home/tyrellshawn/.config/fish/config.fish\` - Updated bun PATH configuration and added cleanup for old opencode installations"
