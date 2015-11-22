GIT_COMMIT=$(shell git rev-parse HEAD)
ENVOY_VERSION=$(shell ./version)
COMPILE_DATE=$(shell date -u +%Y%m%d.%H%M%S)
BUILD_FLAGS=-X main.CompileDate=$(COMPILE_DATE) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(ENVOY_VERSION)

all: build

deps:
	go get github.com/spf13/cobra
	go get github.com/PagerDuty/godspeed
	go get github.com/go-chef/chef
	go get github.com/davecgh/go-spew/spew

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
