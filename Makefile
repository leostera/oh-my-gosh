.PHONY: all ci setup lint benchmark build test

PREFIX ?= /usr/local
GO = $(shell which go)

all: lint build test benchmark

ci: setup all

setup:
	go get -u github.com/alecthomas/gometalinter
	$(shell echo $$GOPATH/bin/gometalinter) --install

lint:
	$(shell echo $$GOPATH/bin/gometalinter) @.gometalinter
	$(GO) vet

benchmark:
	$(GO) test -bench -test.v ./...

build:
	$(GO) build

test:
	$(GO) test -test.v ./...
