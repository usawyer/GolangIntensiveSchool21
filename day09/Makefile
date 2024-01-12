BUILD=${CURDIR}/build
PACKAGE=day09

ex00:
	go build -o ${BUILD}/sleepsort ${PACKAGE}/cmd/ex00
	${BUILD}/sleepsort

ex01:
	go build -o ${BUILD}/crawler ${PACKAGE}/cmd/ex01
	${BUILD}/crawler

ex02:
	go build -o ${BUILD}/fanin ${PACKAGE}/cmd/ex02
	${BUILD}/fanin

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

clean:
	rm -rf ${BUILD}
