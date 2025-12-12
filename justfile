# rhi Justfile

# just --set version v0.3.0
version := "v0.3.0"

# Default recipe
default: build

# Build the binary for the current architecture
build:
    @echo "Building rhi..."
    go build -o bin/rhi ./cmd/rhi

# Run the REPL
run *ARGS:
    go run ./cmd/rhi {{ARGS}}

# Run all tests with the custom test runner flag
test-all:
    go test ./... -v

# Run all tests with the custom test runner flag
test-file PATH:
    go run ./cmd/rhi -test {{PATH}}

# Run the integration tests with your custom flag
test-integration:
    go run ./cmd/rhi -test tests/*.rh

# Compile the ANTLR grammar (wraps go generate)
grammar:
    go generate ./internal/grammar/...

# Create release artifacts for all platforms
release:
    @echo "Building Release {{version}}..."
    mkdir -p dist
    # Linux
    GOOS=linux   GOARCH=amd64 go build -ldflags="-s -w" -o dist/rhi-{{version}}-linux-amd64     ./cmd/rhi
    # Windows
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/rhi-{{version}}-windows-amd64.exe ./cmd/rhi
    # Mac (Intel & Silicon)
    GOOS=darwin  GOARCH=amd64 go build -ldflags="-s -w" -o dist/rhi-{{version}}-darwin-amd64     ./cmd/rhi
    GOOS=darwin  GOARCH=arm64 go build -ldflags="-s -w" -o dist/rhi-{{version}}-darwin-arm64      ./cmd/rhi
    @echo "Done! Artifacts in ./dist"

# Clean up build artifacts
clean:
    rm -rf bin dist