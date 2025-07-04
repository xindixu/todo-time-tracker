.PHONY: build build-server build-cli test clean lint fmt install-tools run run-server run-cli proto-gen help

# Binary names
CLI_BINARY_NAME=ttt
SERVER_BINARY_NAME=ttt-server
CLI_BINARY_PATH=./bin/$(CLI_BINARY_NAME)
SERVER_BINARY_PATH=./bin/$(SERVER_BINARY_NAME)

# Generate protobuf code
proto-gen:
	@echo "Generating protobuf code..."
	mkdir -p proto/go
	protoc --go_out=. --go-grpc_out=. --go_opt=module=todo-time-tracker --go-grpc_opt=module=todo-time-tracker --proto_path=. proto/*.proto

# SQL DB
SQL_DB_SUPER_USER=postgres
SQL_DB_USER=tttuser
SQL_DB_NAME=tttdb
SQL_DB_URL=postgres://$(SQL_DB_USER)@localhost:5432/$(SQL_DB_NAME)?sslmode=disable

sqldb-create:
	@echo "Creating database..."
	createdb $(SQL_DB_NAME)
	createuser $(SQL_DB_USER)
	psql -h localhost -p 5432 -U $(SQL_DB_SUPER_USER) -d $(SQL_DB_NAME) -c "CREATE EXTENSION IF NOT EXISTS btree_gist;"

sqldb-connect:
	@echo "Connecting to database..."
	psql -h localhost -p 5432 -U $(SQL_DB_USER) -d $(SQL_DB_NAME)

sqldb-migrate:
	@echo "Migrating database..."
	SQL_DB_URL=$(SQL_DB_URL) go run cmd/db/main.go

sqldb-drop:
	@echo "Dropping database..."
	dropdb $(SQL_DB_NAME)
	dropuser $(SQL_DB_USER)

# Graph DB
GRAPH_DB_NAME=tttdb
GRAPH_DB_USER=neo4j
GRAPH_DB_PASSWORD=password
GRAPH_DB_URL=neo4j://127.0.0.1:7687

graphdb-create:
	@echo "Creating graph database..."
	echo "CREATE DATABASE ${GRAPH_DB_NAME} IF NOT EXISTS;" | cypher-shell -u neo4j -p password -a neo4j://127.0.0.1:7687 -d system

graphdb-drop:
	@echo "Dropping graph database..."
	echo "DROP DATABASE ${GRAPH_DB_NAME} IF EXISTS;" | cypher-shell -u neo4j -p password -a neo4j://127.0.0.1:7687 -d system

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
server-build:
	@echo "Building $(SERVER_BINARY_NAME)..."
	@mkdir -p bin
	go build -o $(SERVER_BINARY_PATH) ./cmd/server

server-run: server-build
	@echo "Running $(SERVER_BINARY_NAME)..."
	SQL_DB_URL=$(SQL_DB_URL) GRAPH_DB_URI=$(GRAPH_DB_URL) GRAPH_DB_NAME=$(GRAPH_DB_NAME) GRAPH_DB_USER=$(GRAPH_DB_USER) GRAPH_DB_PASSWORD=$(GRAPH_DB_PASSWORD) $(SERVER_BINARY_PATH)

# Docker
server-run-docker:
	@echo "Building and running server with Docker..."
	cd cmd/server && docker compose up --build

server-stop-docker:
	@echo "Stopping Docker containers..."
	cd cmd/server && docker compose down

server-clean-docker:
	@echo "Cleaning Docker containers and images..."
	cd cmd/server && docker compose down --volumes --remove-orphans
	docker system prune -f

server-logs-docker:
	@echo "Showing Docker logs..."
	cd cmd/server && docker compose logs -f

server-restart-docker: server-stop-docker server-run-docker

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
	go vet ./...

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