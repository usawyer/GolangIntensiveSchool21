BUILD=${CURDIR}/build
PACKAGE=bakers

all:

build_ex00:
	go build -o ${BUILD}/readDB ${PACKAGE}/cmd/ex00

run_ex00:
	go run ${PACKAGE}/cmd/ex00

test_ex00: build_ex00
	${BUILD}/readDb -f test/original_database.xml
	${BUILD}/readDb -f test/stolen_database.json

#ex01:
#
#run_ex01:
#
#test_ex01:
#
#ex02:
#
#run_ex02:
#
#test_ex02:

clean:
	rm -rf ${BUILD}/readDb