BUILD=${CURDIR}/build
PACKAGE=cli

all: test_ex00 test_ex01 test_ex02

ex00:
	go build -o ${BUILD}/myFind ${PACKAGE}/cmd/ex00

test_ex00: ex00
	${BUILD}/myFind -f -ext 'go' internal/domain/
	${BUILD}/myFind -d .
	find . -type d

ex01:
	go build -o ${BUILD}/myWc ${PACKAGE}/cmd/ex01

test_ex01: ex01
	${BUILD}/myWc -l test/input.txt test/input2.txt test/input3.txt
	${BUILD}/myWc -m test/input.txt test/input2.txt test/input3.txt
	${BUILD}/myWc -w test/input.txt test/input2.txt test/input3.txt

ex02:
	go build -o ${BUILD}/myXargs ${PACKAGE}/cmd/ex02

test_ex02: ex00 ex01 ex02
	echo 'one two three' | ${BUILD}/myXargs mkdir
	echo 'one two three' | ${BUILD}/myXargs rm -rf
	echo -e "internal/domain/\ncmd/" | ${BUILD}/myXargs ls -la
	echo -e "internal/domain/\ncmd/" | xargs ls -la
	${BUILD}/myFind -f -ext 'go' internal/domain/ | ${BUILD}/myXargs ${BUILD}/myWc -l

ex03:
	go build -o ${BUILD}/myRotate ${PACKAGE}/cmd/ex03

test_ex03: ex03
	${BUILD}/myRotate -a test/ test/input.txt internal/domain/rotate.go

clean:
	rm -rf ${BUILD}
