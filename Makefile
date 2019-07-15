PROTO_FILES := $(shell find ./proto -name *.proto)
PROTO_GEN_DIR := proto/gen/
GO_OUT := $(patsubst ./proto/%.proto,./proto/gen/%.pb.go,$(PROTO_FILES))

setup-lint:
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports

lint:
	go vet
	golint -set_exit_status ./...

fmt: lint
	goimports -w ./

test: lint
	ENV=test go test -cover -race ./...

gen-proto: $(GO_OUT)

lint-proto:
	protolint proto

fmt-proto:
	echo $(PROTO_FILES) | xargs clang-format-8 -i

proto/gen/%.pb.go: proto/%.proto
	protoc -I . $< --go_out=plugins=grpc:$(GOPATH)/src

.PHONY: setup-lint lint fmt test gen-proto lint-proto fmt-proto
