MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
ROOT := $(dir $(MKFILE_PATH))
GOBIN ?= $(ROOT)tools/bin
BIN_PATH ?= $(ROOT)/bin
LINTER_NAME := golangci-lint
LINTER_VERSION := v1.51.2

all: build

.PHONY: build
build: buildexample

.PHONY: buildexample
buildexample:
	go build -o $(BIN_PATH)/example $(ROOT)/cmd/example

.PHONY: test
test: lint
	go test ./...

.PHONY: runexample
runexample: buildexample
	$(BIN_PATH)/example

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: install-linter
install-linter:
	if [ ! -f $(GOBIN)/$(LINTER_VERSION)/$(LINTER_NAME) ]; then \
		echo INSTALLING $(GOBIN)/$(LINTER_VERSION)/$(LINTER_NAME) $(LINTER_VERSION) ; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN)/$(LINTER_VERSION) $(LINTER_VERSION) ; \
		echo DONE ; \
	fi

.PHONY: lint
lint: install-linter
	$(GOBIN)/$(LINTER_VERSION)/$(LINTER_NAME) run --config .golangci.yml
