# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Test Commands

```bash
# Build for current platform
make build

# Install to GOPATH/bin
make install

# Run all tests
make test

# Run tests for specific package
go test ./pkg/config -v

# Run tests with coverage
go test -cover ./...

# Cross-platform build
make release

# Code quality
make fmt      # Format code
make tidy     # Run go mod tidy
make lint     # Run golangci-lint
```

## Architecture Overview

ccconfig is a Go CLI tool for backing up/restoring Claude Code configurations. It has a three-layer architecture:

1. **CLI Layer** (`cmd/`): Cobra commands that parse flags and orchestrate operations
2. **Business Layer** (`pkg/backup/`, `pkg/restore/`, `pkg/cache/`): Core logic for each feature
3. **Support Layer** (`pkg/config`, `pkg/git`, `pkg/i18n`, `pkg/ui`): Cross-cutting utilities

### Key Architecture Patterns

**Configuration Merge**: The system merges configuration from three sources (priority order):
1. Command-line flags (`--repo`, `--lang`, etc.)
2. Config file (`~/.ccconfig.yaml`)
3. Default values

This is handled in `pkg/config/config.go` via viper.

**Internationalization (i18n)**: All user-facing messages use `i18n.T(msgID, templateData)`. The system auto-detects language from:
- Command-line flag `--lang` (highest priority)
- Config file `lang` field
- Environment `LANG`/`LC_ALL`
- Defaults to English

Translation files are embedded via `//go:embed` in `pkg/i18n/i18n.go`.

**Sensitive Data Handling**: API Tokens are NEVER stored in Git:
- Backup: removes `ANTHROPIC_AUTH_TOKEN` from settings.json
- Restore: checks env var, prompts user, or preserves existing token

### Adding a New Command

1. Create `cmd/<name>.go` with a cobra.Command
2. Call business logic from `pkg/` packages
3. Register via `init()` using `GetRootCommand().AddCommand(<name>Cmd)`

### Adding a New Language

1. Create `pkg/i18n/<lang>.yaml` with translations
2. Language code is auto-detected from `LANG` env (e.g., `zh_CN.UTF-8` → `zh`)

## Critical Implementation Details

**i18n.T Signature**: Always call as `i18n.T("msg.id", nil)` for no-params messages or `i18n.T("msg.id", map[string]interface{}{"Key": val})` for template messages.

**File Copying**: Use `backup.CopyFile(src, dst)` from `pkg/backup/commands.go` - it's exported to avoid duplication across packages.

**Git Operations**: All git calls are in `pkg/git/git.go`. The system gracefully handles non-git repos (warnings, continues).

**Error Handling Strategy**:
- Source file missing → Warning, skip, exit 0
- Target permission denied → Error, exit 1
- Git operation fails → Warning, continue
- JSON parsing fails → Error, exit 1
