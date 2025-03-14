.PHONY: build build-all clean release-files

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BINARY_NAME=jc-aws-group-reconciler

# Build the binary for the current platform
build:
	mkdir -p dist
	go build -o dist/${BINARY_NAME} .

# Build the binary for multiple platforms
build-all:
	mkdir -p dist
	# Windows
	GOOS=windows GOARCH=amd64 go build -o dist/${BINARY_NAME}-${VERSION}-windows-amd64.exe .
	# macOS
	GOOS=darwin GOARCH=amd64 go build -o dist/${BINARY_NAME}-${VERSION}-macos-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o dist/${BINARY_NAME}-${VERSION}-macos-arm64 .
	# Linux
	GOOS=linux GOARCH=amd64 go build -o dist/${BINARY_NAME}-${VERSION}-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dist/${BINARY_NAME}-${VERSION}-linux-arm64 .

# Create zip files for each binary
release-files: build-all
	cd dist && \
	sha256sum ${BINARY_NAME}-${VERSION}-* > SHA256SUMS.txt && \
	for f in ${BINARY_NAME}-${VERSION}-* ; do \
		zip -m $$f.zip $$f ; \
	done

# Clean build artifacts
clean:
	rm -rf dist
