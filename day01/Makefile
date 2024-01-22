BUILD=${CURDIR}/build
PACKAGE=bakers

all: test_ex00 test_ex01 test_ex02

ex00:
	go build -o ${BUILD}/readDB ${PACKAGE}/cmd/ex00

test_ex00: ex00
	${BUILD}/readDb -f test/original_database.xml
	${BUILD}/readDb -f test/stolen_database.json

ex01:
	go build -o ${BUILD}/compareDB ${PACKAGE}/cmd/ex01

test_ex01: ex01
	${BUILD}/compareDB --old test/original_database.xml --new test/stolen_database.json

ex02:
	go build -o ${BUILD}/compareFS ${PACKAGE}/cmd/ex02

test_ex02: ex02
	${BUILD}/compareFS --old test/snapshot1.txt --new test/snapshot2.txt

clean:
	rm -rf ${BUILD}
