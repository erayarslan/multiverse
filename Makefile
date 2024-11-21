.PHONY: default

default: init

PROTOBUF_INSTALL_CMD = brew install protobuf
LINT_CMD = golangci-lint run -c .golangci.yml --timeout=5m -v

ifeq ($(OS),Windows_NT)
    PROTOBUF_INSTALL_CMD = choco install protoc
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		PROTOBUF_INSTALL_CMD = sudo apt install -y protobuf-compiler
	endif
endif

init:
	$(PROTOBUF_INSTALL_CMD)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.27.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

lint:
	$(LINT_CMD)

pre-commit:
	go mod tidy
	fieldalignment -fix ./...
	$(LINT_CMD) --fix

proto:
	protoc --go_out=. --go-grpc_out=. --experimental_allow_proto3_optional agent/agent.proto api/api.proto cluster/cluster.proto multipass/multipass.proto

build:
	go build cmd/main.go

test:
	./main -master -worker & echo $$! > pid && sleep 5 && ./main -client -list && kill `cat pid` && rm pid