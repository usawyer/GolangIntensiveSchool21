BUILD=${CURDIR}/build
PACKAGE=search

build_ex00:
	go build -o ${BUILD}/inserter ${PACKAGE}/cmd/ex00

run_ex00: build_ex00
	docker-compose up -d
	@sleep 10
	curl -XPUT "http://localhost:9200/places"
	@sleep 2
	curl -XPUT 'http://localhost:9200/places/_mapping' -H "Content-Type: application/json" -d @api/schema.json
	@sleep 2
	curl -s -XGET "http://localhost:9200/places"
	@sleep 2
	${BUILD}/inserter
	@sleep 5
	curl -s -XGET "http://localhost:9200/places/_doc/1"

build_ex01:
	go build -o ${BUILD}/simple_interface ${PACKAGE}/cmd/ex01

run_ex01: build_ex01
	${BUILD}/simple_interface

clean:
	rm -rf ${BUILD}/inserter
	rm -rf ${BUILD}/simple_interface
	#curl -XDELETE "http://localhost:9200/places/"
	#docker stop day03-elasticsearch
