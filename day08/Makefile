BUILD=${CURDIR}/build
PACKAGE=day08

.PHONY: test

ex01:
	go build -o ${BUILD}/plant_printer ${PACKAGE}/cmd/ex01
	${BUILD}/plant_printer

test:
	go test -v ./test/...

clean:
	rm -rf ${BUILD}
