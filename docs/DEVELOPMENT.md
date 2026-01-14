# Development Guide

This guide is for contributors who want to help improve ccconfig.

---

## Getting Started

### Prerequisites

- Go 1.21+ installed
- Git installed
- Basic knowledge of Go

### Setup

```bash
# Clone the repository
git clone https://github.com/jiangtao/cc-config.git
cd ccconfig

# Download dependencies
go mod download

# Build
make build

# Install to your system
make install
```

---

## Project Structure

```
ccconfig/
├── cmd/              # CLI commands
│   ├── backup.go     # backup command
│   ├── restore.go    # restore command
│   ├── cache.go      # cache command
│   └── init.go       # init command
├── pkg/              # Core packages
│   ├── backup/       # Backup logic
│   ├── restore/      # Restore logic
│   ├── config/       # Configuration
│   ├── git/          # Git operations
│   ├── i18n/         # Translations
│   └── ui/           # Colors & output
├── web/              # Landing page
└── install.sh        # Installation script
```

---

## Development Workflow

### Make Changes

```bash
# Edit code
vim cmd/backup.go

# Run locally
go run . backup --repo ~/test-config

# Format code
make fmt

# Run linter
make lint

# Run tests
make test
```

### Common Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the binary |
| `make install` | Install to system |
| `make test` | Run all tests |
| `make fmt` | Format code |
| `make lint` | Run linter |
| `make tidy` | Clean up dependencies |

---

## Adding Features

### Add a New Command

```go
// cmd/mycommand.go
package cmd

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description of my command",
    RunE:  runMyCommand,
}

func init() {
    GetRootCommand().AddCommand(myCmd)
}

func runMyCommand(cmd *cobra.Command, args []string) error {
    // Your code here
    return nil
}
```

### Add a New Language

1. Create `pkg/i18n/{lang}.yaml`
2. Add translations

```yaml
# pkg/i18n/fr.yaml
my_message:
  title: "My Title"
  description: "My Description"
```

3. Use in code:

```go
i18n.T("my_message.title", nil)
```

---

## Testing

### Write Tests

```go
func TestMyFunction(t *testing.T) {
    input := "test"
    expected := "expected output"

    result := MyFunction(input)

    if result != expected {
        t.Errorf("expected %v, got %v", expected, result)
    }
}
```

### Run Tests

```bash
# Run all tests
make test

# Run specific package
go test ./pkg/backup -v

# Run with coverage
go test -cover ./...
```

---

## Releasing

### Create a Release

```bash
# Create a tag
git tag -a v1.0.0 -m "Release v1.0.0"

# Push the tag
git push origin v1.0.0
```

GitHub Actions will automatically:
- Run tests
- Build binaries for all platforms
- Create a GitHub Release

---

## Contributing

We welcome contributions!

1. Fork the repository
2. Create a branch
   ```bash
   git checkout -b feature/my-feature
   ```
3. Commit your changes
   ```bash
   git commit -m "Add my feature"
   ```
4. Push to your fork
   ```bash
   git push origin feature/my-feature
   ```
5. Create a Pull Request

---

## Tips

### Code Style

- Use `gofmt` to format code
- Follow Go conventions
- Add comments for exported functions

### Debugging

```bash
# Run with verbose output
go run . backup --repo ~/test-config --verbose

# Use Delve debugger
dlv debug ./... -- backup
```

---

## Need Help?

- Create an issue: [https://github.com/jiangtao/cc-config/issues](https://github.com/jiangtao/cc-config/issues)
- Check existing issues: [https://github.com/jiangtao/cc-config/issues](https://github.com/jiangtao/cc-config/issues)

---

## Other Languages

- **中文文档:** [DEVELOPMENT-zh.md](DEVELOPMENT-zh.md)
