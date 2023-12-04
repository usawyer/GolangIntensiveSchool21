package main

import (
	"flag"
	"fmt"
	"gRPC/api/pb"
	"gRPC/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 8888, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	transmitterServer := service.NewTransmitterServer()
	grpcServer := grpc.NewServer()
	pb.RegisterTransmitterServer(grpcServer, transmitterServer)

	address := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
