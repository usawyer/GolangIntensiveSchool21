#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen:
	protoc \
	  --go_out=api \
	  --proto_path=api \
	  --go-grpc_out=api \
	  api/*.proto

pre-build:
	go mod tidy

server-build:pre-build
	go build cmd/server/server.go

server:pre-build
	go run cmd/server/server.go --host localhost --port 8888

client:pre-build
	go run cmd/client/client.go --host localhost --port 8888 --k 2.0 --size 100

client-build:pre-build
	go build cmd/client/client.go

compose-up:
	docker-compose up

clean:
	rm api/pb/*

lint:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -w .
	gofmt -s -w .
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --out-format colored-line-number -v
