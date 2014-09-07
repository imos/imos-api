export GOPATH = $(shell pwd | sed -e 's/\/src\/.*$$//g')
export TARGET = $(shell pwd | sed -e 's/^.*\/src\///g')
HTTP_PORT = :8080

test: build
	go test -v
.PHONY: test

benchmark: build
	go test -v -bench .
.PHONY: benchmark

run: build
	bash "${GOPATH}/src/github.com/imos/imosrpc/server.sh" \
	     "${TARGET}" --imosrpc-http=$(HTTP_PORT)
.PHONY: build

build: get
	go build
.PHONY: build

get: version
	go get
.PHONY: get

version:
	@go version
.PHONY: version

format:
	gofmt -w ./
.PHONY: format

document:
	godoc -http=:6060

info:
	@echo "GOPATH=$${GOPATH}"
.PHONY: info
