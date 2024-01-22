BUILD=${CURDIR}/build
PACKAGE=search

ex00:
	go build -o ${BUILD}/inserter ${PACKAGE}/cmd/ex00

run_ex00: ex00
	docker-compose up -d
	@sleep 20
	curl -XPUT "http://localhost:9200/places"
	@sleep 2
	curl -XPUT 'http://localhost:9200/places/_mapping' -H "Content-Type: application/json" -d @api/schema.json
	@sleep 2
	curl -s -XGET "http://localhost:9200/places"
	@sleep 2
	${BUILD}/inserter
	@sleep 5
	curl -s -XGET "http://localhost:9200/places/_doc/1"

ex01:
	go build -o ${BUILD}/simple_interface ${PACKAGE}/cmd/ex01

run_ex01: ex01
	${BUILD}/simple_interface

ex02:
	go build -o ${BUILD}/proper_API ${PACKAGE}/cmd/ex02

run_ex02: ex02
	${BUILD}/proper_API

ex03:
	go build -o ${BUILD}/closest_restaurants ${PACKAGE}/cmd/ex03

run_ex03: ex03
	${BUILD}/closest_restaurants

ex04:
	go build -o ${BUILD}/jwt ${PACKAGE}/cmd/ex04

run_ex04: ex04
	${BUILD}/jwt
	curl -s -XGET "http://localhost:8888/api/recommend?lat=55.674&lon=37.666" -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.5Mj6ojT5IBSTG-CIN3jCvVndSDnk1qtuTsT3ydcPq-4'

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

clean:
	rm -rf ${BUILD}
	curl -XDELETE "http://localhost:9200/places/"
	docker stop day03-elasticsearch
