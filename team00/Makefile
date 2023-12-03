#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen:
	protoc --proto_path=api --go_out=api --go-grpc_out=api api/*.proto

clean:
	rm api/transmitter/*.go