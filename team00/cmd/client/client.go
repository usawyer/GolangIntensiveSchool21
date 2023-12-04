package main

import (
	"context"
	"flag"
	"fmt"
	"gRPC/api/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	transmitterClient := pb.NewTransmitterClient(conn)

	stream, err := transmitterClient.Transmit(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Transmit failed: %v", err)
	}

	// Чтение и вывод сообщений из стрима
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		fmt.Printf("Received - SessionID: %s, Frequency: %f, Timestamp: %s\n",
			resp.SessionId, resp.Frequency, resp.Timestamp.AsTime().Format(time.RFC3339))
	}
}
