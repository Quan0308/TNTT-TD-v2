# Variables
BINARY_NAME=main-api
CMD_DIR=cmd
BIN_DIR=bin

# Default target
.PHONY: all
all: build

# Run the application
.PHONY: run
run: build
	@./bin/main-api

# Build the application
.PHONY: build
build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_DIR)/main.go

# Run tests
.PHONY: test
test:
	cd main-api && go test ./... -v

# Clean up build artifacts
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

# Format the code
.PHONY: fmt
fmt:
	cd main-api && go fmt ./...

# Install dependencies
.PHONY: deps
deps:
	cd main-api && go mod tidy

# Lint the code
.PHONY: lint
lint:
	cd main-api && golangci-lint run

# Display help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run          Run the application"
	@echo "  make build        Build the binary"
	@echo "  make test         Run all tests"
	@echo "  make clean        Remove build artifacts"
	@echo "  make fmt          Format the code"
	@echo "  make deps         Install and tidy dependencies"
	@echo "  make lint         Run linting (requires golangci-lint)"
