BUILD=${CURDIR}/build
#PACKAGE=cli

all: run_ex00 run_ex01 run_ex02

build_ex00:
	#go build -o ${BUILD}/myFind ${PACKAGE}/cmd/ex00

run_ex00: build_ex00
	docker-compose up -d

clean:
	rm -rf ${BUILD}/myFind
	rm -rf ${BUILD}/myWc
	rm -rf ${BUILD}/myXargs
	rm -rf ${BUILD}/myRotate
