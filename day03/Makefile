BUILD=${CURDIR}/build
#PACKAGE=cli

all: run_ex00 run_ex01 run_ex02

build_ex00:
	#go build -o ${BUILD}/myFind ${PACKAGE}/cmd/ex00

run_ex00: build_ex00
	docker-compose up -d
	curl -XPUT "http://localhost:9200/places"
	curl -XPUT 'http://localhost:9200/places/_mapping' -H "Content-Type: application/json" -d @api/schema.json
	curl -s -XGET "http://localhost:9200/places"

	#curl -XPUT 'http://localhost:9200/places/place/_mapping?include_type_name=true' -H "Content-Type: application/json" -d @"api/schema.json"
	#curl -XPUT "${ES_URL}/places/place/_mapping?include_type_name=true" -H "Content-Type: application/json" -d @"${CURDIR}/build/elastic/schema.json"



clean:
	rm -rf ${BUILD}/myFind
	rm -rf ${BUILD}/myWc
	rm -rf ${BUILD}/myXargs
	rm -rf ${BUILD}/myRotate
