.PHONY: default

default: init

init:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.27.0

linter:
	fieldalignment -fix ./...
	golangci-lint run -c .golangci.yml --timeout=5m -v --fix

tidy:
	go mod tidy