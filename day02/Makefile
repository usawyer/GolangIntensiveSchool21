BUILD=${CURDIR}/build
PACKAGE=cli

all: run_ex00 run_ex01

build_ex00:
	go build -o ${BUILD}/myFind ${PACKAGE}/cmd/ex00

run_ex00: build_ex00
	${BUILD}/myFind -f -ext 'go' internal/domain/
	${BUILD}/myFind -d .
	#find . -type d

build_ex01:
	go build -o ${BUILD}/myWc ${PACKAGE}/cmd/ex01

run_ex01: build_ex01
	${BUILD}/myWc -l test/input.txt test/input2.txt test/input3.txt
	${BUILD}/myWc -m test/input.txt test/input2.txt test/input3.txt
	${BUILD}/myWc -w test/input.txt test/input2.txt test/input3.txt

build_ex02:
	go build -o ${BUILD}/myXargs ${PACKAGE}/cmd/ex02

run_ex02: build_ex00 build_ex01 build_ex02
	echo 'one two three' | ${BUILD}/myXargs mkdir
	#echo 'one two three' | ${BUILD}/myXargs rm -rf
	echo -e "internal/domain/\ncmd/" | ${BUILD}/myXargs ls -la
	echo -e "internal/domain/\ncmd/" | xargs ls -la
	${BUILD}/myFind -f -ext 'go' internal/domain/ | ${BUILD}/myXargs ${BUILD}/myWc -l

clean:
	rm -rf ${BUILD}/myFind
	rm -rf ${BUILD}/myWc
	rm -rf ${BUILD}/myXargs
