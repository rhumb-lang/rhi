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

# Run resolver system tests
test-resolvers:
    go run ./cmd/rhi -test tests/resolvers/resolver_suite.rh
    go run ./cmd/rhi -test tests/fixtures/resolver_tests/src/-/+alias_test.rh
    # Cycle test expects failure, checking return code isn't enough in Just, but it runs.
    go run ./cmd/rhi -test tests/fixtures/resolver_tests/src/-/+cycle_test.rh
    RHUMB_LIB=./tests/fixtures/stdlib go run ./cmd/rhi -test tests/resolvers/stdlib_emoji_test.rh

# Run integrity check tests
test-anchor-validation: build
    @echo "Setting up fixture..."
    mkdir -p tests/integrity_fixture
    @echo 'project:' > tests/integrity_fixture/project@.rhy
    @echo '  📂: src' >> tests/integrity_fixture/project@.rhy
    @echo "'-':" >> tests/integrity_fixture/project@.rhy
    @echo '  test.txt: "text/plain ___"' >> tests/integrity_fixture/project@.rhy
    @echo "Hello World" > tests/integrity_fixture/test.txt
    @echo '1' > tests/integrity_fixture/main.rh

    @echo "Test 1: Auto-fix empty anchor (Expected: Success + File Change)"
    ./bin/rhi -sha256 tests/integrity_fixture/main.rh
    grep -q "sha256:" tests/integrity_fixture/project@.rhy

    @echo "Test 2: Verify Correct Anchor (Expected: Success)"
    ./bin/rhi -sha256 tests/integrity_fixture/main.rh

    @echo "Test 3: Detect Invalid Anchor (Expected: Failure)"
    sed -i 's/sha256:[a-f0-9]*/sha256:BAD/' tests/integrity_fixture/project@.rhy
    ! ./bin/rhi -sha256 tests/integrity_fixture/main.rh

    @echo "Cleanup..."
    rm -rf tests/integrity_fixture
    @echo "Anchor Validation Passed!"

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