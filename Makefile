.PHONY: build test clean

BINARY_NAME=github-vikunja-sync
BUILD_DIR=build
VERSION=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-X main.version=${VERSION}"

build:
	@echo "Building ${BINARY_NAME}..."
	@mkdir -p ${BUILD_DIR}
	@go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ./main.go

test:
	@go test -v ./...

clean:
	@rm -rf ${BUILD_DIR}
