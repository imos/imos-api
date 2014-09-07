export GOPATH = $(shell pwd)
TARGET = \
	api/info \
	config
HTTP_PORT = :8080

test: build
	go test -v $(TARGET)
.PHONY: test

benchmark: build
	go test -v -bench $(TARGET)
.PHONY: benchmark

run: build
	bash "${GOPATH}/src/github.com/imos/imosrpc/server.sh" \
	     "api" --imosrpc-http=$(HTTP_PORT)
.PHONY: build

appengine: build
	dev_appserver.py \
	    --host 0.0.0.0 --port 8080 \
	    --admin_host 0.0.0.0 --admin_port 8000 \
	    --skip_sdk_update_check true \
	    src
.PHONY: appengine

build: get
	go build $(TARGET)
.PHONY: build

get: version
	go get $(TARGET)
.PHONY: get

version:
	@go version
.PHONY: version

format:
	gofmt -w $(TARGET)
.PHONY: format

document:
	godoc -http=:6060

info:
	@echo "GOPATH=$${GOPATH}"
.PHONY: info
