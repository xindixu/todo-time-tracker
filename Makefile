.PHONY: build test clean lint fmt install-tools run help test-hello

# Binary name
BINARY_NAME=ttt
BINARY_PATH=./bin/$(BINARY_NAME)

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	go build -o $(BINARY_PATH) ./cmd/ttt

# Build and run hello world test
test-hello:
	@echo "Running hello world test..."
	go run test_hello.go

# Run the application
run: build
	$(BINARY_PATH)

# Test the application
test:
	@echo "Running tests..."
	go test -v ./...

# Test with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -f test_hello

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Running linter..."
	golangci-lint run

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME) to GOPATH/bin..."
	cp $(BINARY_PATH) $(GOPATH)/bin/

# Show help
help:
	@echo "Available targets:"
	@echo "  build         Build the application"
	@echo "  run           Build and run the application"
	@echo "  test          Run tests"
	@echo "  test-hello    Run hello world test"
	@echo "  test-coverage Run tests with coverage report"
	@echo "  clean         Clean build artifacts"
	@echo "  fmt           Format code"
	@echo "  lint          Run linter"
	@echo "  install-tools Install development tools"
	@echo "  install       Install binary to GOPATH/bin"
	@echo "  help          Show this help message"

# Default target
.DEFAULT_GOAL := help