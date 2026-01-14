# Architecture

## Overview

ccconfig uses a simple three-layer architecture:

1. **CLI Layer** (`cmd/`) - Command-line interface
2. **Business Layer** (`pkg/backup/`, `pkg/restore/`) - Core logic
3. **Support Layer** (`pkg/config`, `pkg/git`, `pkg/i18n`, `pkg/ui`) - Utilities

```
┌─────────────────────────────────────────────────────────┐
│                         CLI Layer                        │
│  User commands: backup, restore, cache, init             │
└─────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                      Business Layer                      │
│  Backup and restore logic for settings, commands, etc.   │
└─────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                      Support Layer                       │
│  Config, Git, i18n, UI helpers                          │
└─────────────────────────────────────────────────────────┘
```

## Key Packages

| Package | Purpose |
|---------|---------|
| `pkg/backup/` | Backup settings, commands, skills, projects |
| `pkg/restore/` | Restore from backup |
| `pkg/config/` | Load and merge config from flags + files |
| `pkg/git/` | Git operations (init, add, commit, push) |
| `pkg/i18n/` | Multi-language support (English/中文) |

## How It Works

### Backup Flow

```
User runs: ccconfig backup
    │
    ▼
Load config (flags + ~/.ccconfig.yaml)
    │
    ▼
Backup settings → commands → skills → projects
    │
    ▼
Git add + commit (optional: push)
    │
    ▼
Done! ✅
```

### Restore Flow

```
User runs: ccconfig restore
    │
    ▼
Git pull (get latest configs)
    │
    ▼
Restore settings → commands → skills → projects
    │
    ▼
Prompt for API token (if needed)
    │
    ▼
Done! ✅
```

## For Developers

### Add a new command

1. Create file in `cmd/`
2. Use Cobra library
3. Register in `init()`

### Add a new language

1. Create `pkg/i18n/{lang}.yaml`
2. Add translations

For more details, see [DEVELOPMENT.md](DEVELOPMENT.md).

---

## Other Languages

- **中文文档:** [ARCHITECTURE-zh.md](ARCHITECTURE-zh.md)
