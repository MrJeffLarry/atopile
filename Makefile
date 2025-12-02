# Makefile for atopile Go version

# Build variables
BINARY_NAME=ato
OUTPUT_DIR=bin
MAIN_PACKAGE=./cmd/ato

# Version information
VERSION?=dev
GIT_COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE?=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Build flags
LDFLAGS=-ldflags "\
	-X 'github.com/atopile/atopile/internal/version.Version=$(VERSION)' \
	-X 'github.com/atopile/atopile/internal/version.GitCommit=$(GIT_COMMIT)' \
	-X 'github.com/atopile/atopile/internal/version.BuildDate=$(BUILD_DATE)'"

.PHONY: all build test clean install deps fmt vet lint help

all: test build

## build: Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(OUTPUT_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)"

## test: Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v -race -coverprofile=coverage.txt -covermode=atomic ./...

## test-short: Run short tests only
test-short:
	@echo "Running short tests..."
	$(GOTEST) -short ./...

## coverage: Generate coverage report
coverage: test
	@echo "Generating coverage report..."
	$(GOCMD) tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(OUTPUT_DIR)
	rm -f coverage.txt coverage.html

## install: Install the binary to GOPATH/bin
install:
	@echo "Installing..."
	$(GOCMD) install $(LDFLAGS) $(MAIN_PACKAGE)

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

## fmt: Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

## vet: Run go vet
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

## lint: Run linters (requires golangci-lint)
lint:
	@echo "Running linters..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install from https://golangci-lint.run/"; \
	fi

## check: Run fmt, vet, and test
check: fmt vet test

## run: Build and run the CLI
run: build
	@./$(OUTPUT_DIR)/$(BINARY_NAME)

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

# Default target
.DEFAULT_GOAL := help
