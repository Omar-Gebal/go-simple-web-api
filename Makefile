APP_NAME := simple-weather-api
BUILD_DIR := build

# Default target: build the application
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(APP_NAME)

# Clean the build directory
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	@go fmt ./...

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy


# Generate a hashed password
.PHONY: generate-password
generate-password:
	@go run ./cmd/password_generator/main.go

# Display help
.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo
	@echo "Usage:"
	@echo "  make [target]"
	@echo
	@echo "Targets:"
	@echo "  all             - Build the application (default)"
	@echo "  build           - Build the application"
	@echo "  run             - Run the application"
	@echo "  clean           - Clean the build directory"
	@echo "  fmt             - Format the code"
	@echo "  test            - Run tests"
	@echo "  deps            - Install dependencies"
	@echo "  generate-password - Generate a hashed password"
	@echo "  help            - Display this help message"
