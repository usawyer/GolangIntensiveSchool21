.PHONY: test

test:
	go test -v ./test/...

doc:
	go install golang.org/x/tools/cmd/godoc@latest
	godoc -http=:8080
	http://localhost:8080/pkg/day07/

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run
