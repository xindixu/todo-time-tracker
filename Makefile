.PHONY: build build-server build-cli test clean lint fmt install-tools run run-server proto-gen help

# Binary names
CLI_BINARY_NAME=ttt
SERVER_BINARY_NAME=ttt-server
CLI_BINARY_PATH=./bin/$(CLI_BINARY_NAME)
SERVER_BINARY_PATH=./bin/$(SERVER_BINARY_NAME)

# Generate protobuf code
proto-gen:
	@echo "Generating protobuf code..."
    protoc --go_out=. --go-grpc_out=. proto/*.proto

# Build the server
build-server:
	@echo "Building $(SERVER_BINARY_NAME)..."
	@mkdir -p bin
	go build -o $(SERVER_BINARY_PATH) ./cmd/server

# Build the CLI client
build-cli:
	@echo "Building $(CLI_BINARY_NAME)..."
	@mkdir -p bin
	go build -o $(CLI_BINARY_PATH) ./cmd/cli

# Build both server and client
build: build-server build-cli

# Build and run hello world test
test-hello:
	@echo "Running hello world test..."
	go run test_hello.go

# Run the server
run-server: build-server
	$(SERVER_BINARY_PATH)

# Run the CLI (you'll need to specify arguments)
run: build-cli
	$(CLI_BINARY_PATH)

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
	@echo "  proto-gen     Generate protobuf Go code"
	@echo "  build         Build both server and CLI"
	@echo "  build-server  Build the gRPC server"
	@echo "  build-cli     Build the CLI client"
	@echo "  run-server    Build and run the server"
	@echo "  run           Build and run the CLI"
	@echo "  test          Run tests"
	@echo "  test-coverage Run tests with coverage report"
	@echo "  clean         Clean build artifacts"
	@echo "  fmt           Format code"
	@echo "  lint          Run linter"
	@echo "  install-tools Install development tools"
	@echo "  install       Install binary to GOPATH/bin"
	@echo "  help          Show this help message"

# Default target
.DEFAULT_GOAL := help