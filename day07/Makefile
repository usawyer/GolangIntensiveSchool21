.PHONY: test

test:
	go test -v ./test/...

doc:
	go install golang.org/x/tools/cmd/godoc@latest
	~/go/bin/godoc -http=:8080

gen:
	@mkdir -p docs
	wget --mirror --convert-links --adjust-extension --page-requisites --no-parent --no-host-directories -P docs/ http://localhost:8080/pkg/day07/internal
	zip -r docs/docs.zip docs
	rm -rf docs/lib docs/pkg

unzip:
	unzip docs/docs.zip

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

clean:
	rm -rf docs
