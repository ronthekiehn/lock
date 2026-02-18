BIN_NAME := lock
BUILD_DIR := build
BIN_PATH := $(BUILD_DIR)/$(BIN_NAME)
PREFIX ?= /usr/local/bin
VERSION ?= $(shell \
	TAG=$$(git describe --tags --exact-match 2>/dev/null || true); \
	COMMIT=$$(git rev-parse --short HEAD 2>/dev/null || echo unknown); \
	DIRTY=$$(test -n "$$(git status --porcelain 2>/dev/null)" && echo ".dirty" || true); \
	if [ -n "$$TAG" ]; then \
		echo "$$TAG"; \
	else \
		echo "0.0.0-dev+$$COMMIT$$DIRTY"; \
	fi \
)
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
