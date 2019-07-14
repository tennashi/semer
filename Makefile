PROTO_FILES := $(shell find ./proto -name *.proto)
PROTO_GEN_DIR := proto/gen/
GO_OUT := $(patsubst ./proto/%.proto,./proto/gen/%.pb.go,$(PROTO_FILES))

gen-proto: $(GO_OUT)

lint-proto:
	protolint proto

fmt-proto:
	echo $(PROTO_FILES) | xargs clang-format-8 -i

proto/gen/%.pb.go: proto/%.proto
	protoc -I . $< --go_out=plugins=grpc:$(GOPATH)/src

.PHONY: gen-proto
