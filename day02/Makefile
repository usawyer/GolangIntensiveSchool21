BUILD=${CURDIR}/build
PACKAGE=cli

all: run_ex00 run_ex01

build_ex00:
	go build -o ${BUILD}/myFind ${PACKAGE}/cmd/ex00

run_ex00: build_ex00
	${BUILD}/myFind -f test/original_database.xml
# указать путь

build_ex01:
	go build -o ${BUILD}/myWc ${PACKAGE}/cmd/ex01

run_ex01: build_ex01
	${BUILD}/myWc -l test/input.txt test/input2.txt test/input3.txt
	${BUILD}/myWc -m test/input.txt test/input2.txt test/input3.txt
	${BUILD}/myWc -w test/input.txt test/input2.txt test/input3.txt

clean:
	rm -rf ${BUILD}/myFind
	rm -rf ${BUILD}/myWc
