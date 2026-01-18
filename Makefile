BIN := orca
BUILD_DIR := build

.PHONY: help build test test-all test-watch lint debug clean

help:
	@echo "make build        - build binary"
	@echo "make test-all     - all tests"
	@echo "make test-watch   - watch tests"
	@echo "make lint         - go vet / staticcheck"
	@echo "make clean        - clean build"

build:
	@scripts/build.sh


test-all:
	@scripts/test.sh ./...

test-watch:
	@scripts/test-watch.sh internal/...

lint:
	@scripts/lint.sh

clean:
	rm -rf $(BUILD_DIR)