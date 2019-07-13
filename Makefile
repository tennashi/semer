PROTO_FILES := $(shell find ./proto -name *.proto)
PROTO_GEN_DIR := proto/gen/
GO_OUT := $(patsubst ./proto/%.proto,./proto/gen/%.pb.go,$(PROTO_FILES))

gen-proto: $(GO_OUT)

lint-proto:
	protolint proto

fmt-proto:
	echo $(PROTO_FILES) | xargs clang-format-8 -i

proto/gen/%.pb.go: proto/%.proto
	protoc -I ./proto $< --go_out=plugins=grpc:./proto/gen

.PHONY: gen-proto
