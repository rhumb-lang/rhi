# rhi Justfile

# Default recipe
default: build

# Build the binary for the current architecture
build:
    @echo "Building rhi..."
    go build -o bin/rhi ./cmd/rhi

# Run the REPL
run:
    go run ./cmd/rhi

# Run all tests with the custom test runner flag
test:
    go test ./... -v

# Run the integration tests with your custom flag
test-integration:
    go run ./cmd/rhi -test tests/*.rh

# Compile the ANTLR grammar (wraps go generate)
grammar:
    go generate ./internal/grammar/...

# Create release artifacts for all platforms
release:
    @echo "Building Release v0.2.0..."
    mkdir -p dist
    # Linux
    GOOS=linux   GOARCH=amd64 go build -ldflags="-s -w" -o dist/rhi-v0.2.0-linux-amd64     ./cmd/rhi
    # Windows
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/rhi-v0.2.0-windows-amd64.exe ./cmd/rhi
    # Mac (Intel & Silicon)
    GOOS=darwin  GOARCH=amd64 go build -ldflags="-s -w" -o dist/rhi-v0.2.0-darwin-amd64     ./cmd/rhi
    GOOS=darwin  GOARCH=arm64 go build -ldflags="-s -w" -o dist/rhi-v0.2.0-darwin-arm64      ./cmd/rhi
    @echo "Done! Artifacts in ./dist"

# Clean up build artifacts
clean:
    rm -rf bin dist