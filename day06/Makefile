BUILD=${CURDIR}/build
PACKAGE=day06

ex00:
	go build -o ${BUILD}/amazing_logo ${PACKAGE}/cmd/ex00
	${BUILD}/amazing_logo --path ${CURDIR}/build
	cp ${BUILD}/amazing_logo.png assets/

ex01:
	docker-compose up

clean:
	rm -rf ${BUILD}
	rm -rf assets/amazing_logo.png
