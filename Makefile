.PHONY: build build-server build-cli test clean lint fmt install-tools run run-server run-cli proto-gen help

# Binary names
CLI_BINARY_NAME=ttt
SERVER_BINARY_NAME=ttt-server
CLI_BINARY_PATH=./bin/$(CLI_BINARY_NAME)
SERVER_BINARY_PATH=./bin/$(SERVER_BINARY_NAME)

# Generate protobuf code
proto-gen:
	@echo "Generating protobuf code..."
    protoc --go_out=. --go-grpc_out=. proto/*.proto


# Database
DATABASE_SUPERUSER=postgres
DATABASE_USER=tttuser
DATABASE_NAME=tttdb
DATABASE_URL=postgres://$(DATABASE_USER)@localhost:5432/$(DATABASE_NAME)?sslmode=disable

db-create:
	@echo "Creating database..."
	createdb $(DATABASE_NAME)
	createuser $(DATABASE_USER)
	psql -h localhost -p 5432 -U $(DATABASE_SUPERUSER) -d $(DATABASE_NAME) -c "CREATE EXTENSION IF NOT EXISTS btree_gist;"

db-connect:
	@echo "Connecting to database..."
	psql -h localhost -p 5432 -U $(DATABASE_USER) -d $(DATABASE_NAME)

db-migrate:
	@echo "Migrating database..."
	DATABASE_URL=$(DATABASE_URL) go run cmd/db/main.go

db-drop:
	@echo "Dropping database..."
	dropdb $(DATABASE_NAME)
	dropuser $(DATABASE_USER)

# CLI
build-cli:
	@echo "Building $(CLI_BINARY_NAME)..."
	@mkdir -p bin
	go build -o $(CLI_BINARY_PATH) ./cmd/cli

run-cli: build-cli
	@echo "Running $(CLI_BINARY_NAME)..."
	$(CLI_BINARY_PATH)

# Build both CLI and server
build: build-cli build-server
	@echo "Built both CLI and server binaries"

# Server
build-server:
	@echo "Building $(SERVER_BINARY_NAME)..."
	@mkdir -p bin
	go build -o $(SERVER_BINARY_PATH) ./cmd/server

run-server: build-server
	@echo "Running $(SERVER_BINARY_NAME)..."
	DATABASE_URL=$(DATABASE_URL) $(SERVER_BINARY_PATH)


# Docker
# run-server-docker:
# 	@echo "Building and running server with Docker..."
# 	cd cmd/server && docker compose up --build

# stop-server-docker:
# 	@echo "Stopping Docker containers..."
# 	cd cmd/server && docker compose down

# clean-docker:
# 	@echo "Cleaning Docker containers and images..."
# 	cd cmd/server && docker compose down --volumes --remove-orphans
# 	docker system prune -f

# logs-docker:
# 	@echo "Showing Docker logs..."
# 	cd cmd/server && docker compose logs -f

# restart-server-docker: stop-server-docker run-server-docker




# # Run the CLI (you'll need to specify arguments)
# run: build-cli
# 	$(CLI_BINARY_PATH)
# Development

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
	@echo "  proto-gen              Generate protobuf Go code"
	@echo "  build                  Build both CLI and server"
	@echo "  build-cli              Build the CLI application"
	@echo "  run-cli                Build and run the CLI application"
	@echo "  build-server           Build the gRPC server (local)"
	@echo "  run-server             Build and run the server (local)"
	@echo "  test                   Run tests"
	@echo "  test-coverage          Run tests with coverage report"
	@echo "  clean                  Clean build artifacts"
	@echo "  fmt                    Format code"
	@echo "  lint                   Run linter"
	@echo "  install-tools          Install development tools"
	@echo "  install                Install binary to GOPATH/bin"
	@echo "  help                   Show this help message"

# Default target
.DEFAULT_GOAL := help