export GO111MODULE=on

SHELL=/bin/bash -o pipefail
PWD = $(shell pwd)
GO ?= go
DOCKER ?= docker
PROTOC ?= protoc
STRIP = strip

.PHONY: all
all: immu immud

.PHONY: immu
immu:
	$(GO) build ./cmd/immu

.PHONY: immud
immud:
	$(GO) build ./cmd/immud

.PHONY: immu-static
immu-static:
	$(GO) build -a -tags netgo -ldflags '${LDFLAGS} -extldflags "-static"' ./cmd/immu

.PHONY: immud-static
immud-static:
	$(GO) build -a -tags netgo -ldflags '${LDFLAGS} -extldflags "-static"' ./cmd/immud

.PHONY: vendor
vendor:
	$(GO) mod vendor

.PHONY: test
test:
	$(GO) vet ./...
	$(GO) test --race ${TEST_FLAGS} ./...

.PHONY: build/codegen
build/codegen:
	$(PROTOC) -I pkg/schema/ pkg/schema/schema.proto --go_out=plugins=grpc:pkg/schema

.PHONY: clean
clean:
	rm -f immu immud bm

.PHONY: nimmu
nimmu:
	$(GO) build -o nimmu ./tools/nimmu 

.PHONY: bm
bm:
	$(GO) build -ldflags '-s -w' ./tools/bm
	$(STRIP) bm

.PHONY: bm/function
bm/function: bm
	./bm function

.PHONY: bm/rpc
bm/rpc: bm
	./bm rpc

.PHONY: bench
bench:
	$(DOCKER) build -t immu_bench -f ./Dockerfile.bench .
	$(DOCKER) run --rm -it immu_bench

.PHONY: tools/comparison/mongodb
tools/comparison/mongodb:
	$(DOCKER) build -t immu_mongodb ./tools/comparison/mongodb
	$(DOCKER) run --rm -it immu_mongodb

.PHONY: tools/comparison/scylladb
tools/comparison/scylladb:
	$(DOCKER) build -t immu_scylladb ./tools/comparison/scylladb
	$(DOCKER) run --rm -it immu_scylladb
