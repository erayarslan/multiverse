.PHONY: default

default: init

PROTOBUF_INSTALL_CMD = brew install protobuf

ifeq ($(OS),Windows_NT)
    PROTOBUF_INSTALL_CMD = choco install protoc
endif

windows:
	go mod tidy
	go build -o main.exe cmd/main.go

init:
	$(PROTOBUF_INSTALL_CMD)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.27.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

proto:
	protoc --go_out=. --go-grpc_out=. agent/agent.proto api/api.proto cluster/cluster.proto multipass/multipass.proto


linter:
	go mod tidy
	fieldalignment -fix ./...
	golangci-lint run -c .golangci.yml --timeout=5m -v --fix

build:
	go build cmd/main.go

test:
	./main -master -worker & echo $$! > pid && sleep 5 && ./main -client -list && kill `cat pid` && rm pid