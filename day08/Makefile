BUILD=${CURDIR}/build
PACKAGE=day08

.PHONY: test

ex00:
	go test -v ./test/ex00/...

ex01:
	go test -v ./test/ex01/...

ex02:
	go build -o ${BUILD}/window ${PACKAGE}/cmd/ex02
	${BUILD}/window

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

clean:
	rm -rf ${BUILD}
