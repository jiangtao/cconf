# Makefile for ccconfig

BINARY_NAME=ccconfig
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Go build flags
GO=go
GOFLAGS=-v

# Platform mappings
PLATFORMS=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64

.PHONY: all build clean test install release help

# Default target
all: build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME) for current platform..."
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BINARY_NAME) .

# Install to GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GO) install $(GOFLAGS) $(LDFLAGS) .

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -rf dist

# Run tests
test:
	$(GO) test -v ./...

# Build for all platforms
release: clean
	@echo "Building release binaries..."
	@$(MAKE) $(PLATFORMS)

# Platform builds
darwin/amd64:
	@echo "Building for darwin/amd64..."
	@mkdir -p dist
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .

darwin/arm64:
	@echo "Building for darwin/arm64..."
	@mkdir -p dist
	GOOS=darwin GOARCH=arm64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .

linux/amd64:
	@echo "Building for linux/amd64..."
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .

linux/arm64:
	@echo "Building for linux/arm64..."
	@mkdir -p dist
	GOOS=linux GOARCH=arm64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 .

windows/amd64:
	@echo "Building for windows/amd64..."
	@mkdir -p dist
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .

# Create release archives
archives: release
	@echo "Creating release archives..."
	@cd dist && \
	for binary in $(BINARY_NAME)-*; do \
		if [[ $$binary == *.exe ]]; then \
			zip $${binary%.exe}.zip $$binary; \
		else \
			tar czf $${binary}.tar.gz $$binary; \
		fi; \
	done

# Run linter
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run: brew install golangci-lint"; \
	fi

# Run formatter
fmt:
	$(GO) fmt ./...

# Run go mod tidy
tidy:
	$(GO) mod tidy

# Show help
help:
	@echo "Available targets:"
	@echo "  all        - Build for current platform (default)"
	@echo "  build      - Build for current platform"
	@echo "  install    - Install to GOPATH/bin"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  release    - Build for all platforms"
	@echo "  archives   - Create release archives"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo "  tidy       - Run go mod tidy"
	@echo "  help       - Show this help"
