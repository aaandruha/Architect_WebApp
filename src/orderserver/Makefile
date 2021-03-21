# This file defines GNU Make targets

GOPATH=$(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

all: build test check

.PHONY: modules
modules:
	go mod tidy

# Builds Go binaries
.PHONY: build
build: modules
	cp -r test_data bin
	$(foreach name,$(APP_CMD_NAMES),bin/run-go-build "cmd/$(name)" "bin/$(name)")

# Removes built Go binaries
.PHONY: clean
clean:
	$(foreach name,$(APP_CMD_NAMES),rm -f "bin/$(name)")

.PHONY: test
test:
	go test ./...

.PHONY: check
check:
	golangci-lint run
