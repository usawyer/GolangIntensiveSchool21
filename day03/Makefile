BUILD=${CURDIR}/build
PACKAGE=search

#all: run_ex00 run_ex01 run_ex02

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

	#curl -XDELETE "http://localhost:9200/places/_doc/1"
	#curl -XDELETE "http://localhost:9200/places/"
	#curl -XPUT 'http://localhost:9200/places/place/_mapping?include_type_name=true' -H "Content-Type: application/json" -d @"api/schema.json"
	#curl -XPUT "${ES_URL}/places/place/_mapping?include_type_name=true" -H "Content-Type: application/json" -d @"${CURDIR}/build/elastic/schema.json"



clean:
	rm -rf ${BUILD}/insert
