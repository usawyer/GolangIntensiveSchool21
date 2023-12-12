.PHONY: test

test:
	go test -v ./test/...

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run
