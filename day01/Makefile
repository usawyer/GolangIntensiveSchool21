BUILD=${CURDIR}/build
PACKAGE=bakers

all: test_ex00 test_ex01 test_ex02

build_ex00:
	go build -o ${BUILD}/readDB ${PACKAGE}/cmd/ex00

run_ex00:
	go run ${PACKAGE}/cmd/ex00

test_ex00: build_ex00
	${BUILD}/readDb -f test/original_database.xml
	${BUILD}/readDb -f test/stolen_database.json

build_ex01:
	go build -o ${BUILD}/compareDB ${PACKAGE}/cmd/ex01

run_ex01:
	go run ${PACKAGE}/cmd/ex01

test_ex01: build_ex01
	${BUILD}/compareDB --old test/original_database.xml --new test/stolen_database.json

build_ex02:
	go build -o ${BUILD}/compareDB ${PACKAGE}/cmd/ex01

run_ex02:
	go run ${PACKAGE}/cmd/ex02

test_ex02: build_ex02
	${BUILD}/compareFS --old test/snapshot1.txt --new test/snapshot2.txt

clean:
	rm -rf ${BUILD}/readDb
	rm -rf ${BUILD}/compareDB
	rm -rf ${BUILD}/compareFS