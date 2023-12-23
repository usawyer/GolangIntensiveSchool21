BUILD=${CURDIR}/build
PACKAGE=day09

ex00:
	go build -o ${BUILD}/sleepsort ${PACKAGE}/cmd/ex00
	${BUILD}/sleepsort

ex02:
	go build -o ${BUILD}/fanin ${PACKAGE}/cmd/ex02
	${BUILD}/fanin

test:
	go test -v ./test/...

clean:
	rm -rf ${BUILD}
