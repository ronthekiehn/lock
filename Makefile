BIN_NAME := lock
BUILD_DIR := build
BIN_PATH := $(BUILD_DIR)/$(BIN_NAME)
PREFIX ?= /usr/local/bin
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -X main.version=$(VERSION)

.PHONY: build install uninstall clean release-local

build:
	mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BIN_PATH) ./cmd/lock

install: build
	if [ -w "$(PREFIX)" ]; then \
		install -m 0755 $(BIN_PATH) "$(PREFIX)/$(BIN_NAME)"; \
	else \
		sudo install -m 0755 $(BIN_PATH) "$(PREFIX)/$(BIN_NAME)"; \
	fi

uninstall:
	if [ -w "$(PREFIX)" ]; then \
		rm -f "$(PREFIX)/$(BIN_NAME)"; \
	else \
		sudo rm -f "$(PREFIX)/$(BIN_NAME)"; \
	fi

clean:
	rm -rf $(BUILD_DIR)

release-local:
	@if ! command -v goreleaser >/dev/null 2>&1; then \
		echo "goreleaser is not installed"; \
		exit 1; \
	fi
	goreleaser build --snapshot --clean
