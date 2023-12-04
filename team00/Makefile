#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen:
	protoc --proto_path=api --go_out=api --go-grpc_out=api api/*.proto

server:
	go run cmd/server/server.go --port 8888

client:
	go run cmd/client/client.go --address localhost:8888

clean:
	rm api/pb/*