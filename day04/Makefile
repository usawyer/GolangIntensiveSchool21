BUILD=${CURDIR}/build
PACKAGE=candies

.PHONY: build

build:
	go build -o ${BUILD}/server ${PACKAGE}/cmd/candies-server
	go build -o ${BUILD}/client ${PACKAGE}/cmd/candies-client

run_ex00: build
	${BUILD}/server --port 3333

test_ex00: build
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": 1}' http://127.0.0.1:3333/buy_candy
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 46, "candyType": "YR", "candyCount": 2}' http://127.0.0.1:3333/buy_candy

run_ex01: build
	${BUILD}/server --scheme https --tls-port 3333 \
	--tls-certificate cert/server/cert.pem --tls-key cert/server/key.pem

test_ex01: build
	${BUILD}/client -k AA -c 2 -m 50
	${BUILD}/client -k DE -c 1 -m 21
	${BUILD}/client -k YR -c 4 -m 13

run_ex02: build
	${BUILD}/server --scheme https --tls-port 3333 \
    --tls-certificate cert/server/cert.pem --tls-key cert/server/key.pem --cow

test_ex02: build
	curl -s --key cert/server/key.pem --cert cert/server/cert.pem --cacert cert/minica.pem \
    -XPOST -H "Content-Type: application/json" -d '{"candyType": "NT", "candyCount": 2, "money": 34}' \
    "https://127.0.0.1:3333/buy_candy"

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

clean:
	rm -rf ${BUILD}/server
	rm -rf ${BUILD}/client
