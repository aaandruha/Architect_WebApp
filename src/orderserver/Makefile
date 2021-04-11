# This file defines GNU Make targets

APP_NAME=orderserver
GOPATH=$(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

all: build #test check

.PHONY: modules
modules:
	go mod tidy

# Builds Go binaries
.PHONY: build
build: modules
	go build -o bin/$(APP_NAME) cmd/orderservice/main.go

# Removes built Go binaries
.PHONY: clean
clean:
	rm -f bin/$(APP_NAME)

.PHONY: test
test:
	go test ./...

.PHONY: check
check:
	golangci-lint run
