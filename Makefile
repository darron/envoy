ENVOY_VERSION="0.4-dev"
GIT_COMMIT=$(shell git rev-parse HEAD)
COMPILE_DATE=$(shell date -u +%Y%m%d.%H%M%S)
BUILD_FLAGS=-X main.CompileDate=$(COMPILE_DATE) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(ENVOY_VERSION)

all: build

deps:
	go get -u github.com/spf13/cobra
	go get -u github.com/PagerDuty/godspeed
	go get -u github.com/go-chef/chef
	go get -u github.com/davecgh/go-spew/spew
	go get -u github.com/zorkian/go-datadog-api
	go get -u github.com/darron/envoy

format:
	gofmt -w .

clean:
	rm -f bin/envoy || true

build: clean
	go build -ldflags "$(BUILD_FLAGS)" -o bin/envoy main.go

gziposx:
	gzip bin/envoy
	mv bin/envoy.gz bin/envoy-$(ENVOY_VERSION)-darwin.gz

linux: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o bin/envoy main.go

gziplinux:
	gzip bin/envoy
	mv bin/envoy.gz bin/envoy-$(ENVOY_VERSION)-linux-amd64.gz

release: clean build gziposx clean linux gziplinux clean
