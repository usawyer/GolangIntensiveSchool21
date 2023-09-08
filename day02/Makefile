BUILD=${CURDIR}/build
PACKAGE=cli

all: test_ex00

build_ex00:
	go build -o ${BUILD}/myFind ${PACKAGE}/cmd/ex00

run_ex00:
	go run ${PACKAGE}/cmd/ex00

test_ex00: build_ex00
	${BUILD}/myFind -f test/original_database.xml

clean:
	rm -rf ${BUILD}/myFind
