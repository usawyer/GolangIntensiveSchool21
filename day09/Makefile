BUILD=${CURDIR}/build
PACKAGE=day09

ex00:
	go build -o ${BUILD}/sleepsort ${PACKAGE}/cmd/ex00
	${BUILD}/sleepsort

test:
	go test -v ./test/...

clean:
	rm -rf ${BUILD}
