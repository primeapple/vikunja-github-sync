.PHONY: build test clean install format fmtcheck vet check

BINARY_NAME=github-vikunja-sync
BUILD_DIR=build
VERSION=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-X main.version=${VERSION}"

build:
	@echo "Building ${BINARY_NAME}..."
	@mkdir -p ${BUILD_DIR}
	@go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ./main.go

test:
	@echo "Running tests..."
	@go test -v ./...

vet:
	@echo "Running lint..."
	@go vet ./...

formatcheck:
	@echo "Checking formatting..."
	@if [ -n "$(shell gofmt -l .)" ]; then \
		echo "Wrong formatting for files:"; \
		gofmt -l .; \
		echo "Run 'make format' to fix formatting issues."; \
		exit 1; \
	fi
	@go mod tidy


format:
	@echo "Formatting code..."
	@gofmt -s -w .
	@go mod tidy

all: format vet test build
	@echo "All checks passed!"

clean:
	@rm -rf ${BUILD_DIR}
