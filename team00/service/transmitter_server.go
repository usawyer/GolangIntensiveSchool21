package service

import (
	"gRPC/api/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"time"
)

type TransmitterServer struct {
	pb.UnimplementedTransmitterServer
}

func NewTransmitterServer() *TransmitterServer {
	return &TransmitterServer{}
}

func (server *TransmitterServer) Transmit(
	req *pb.Request,
	stream pb.Transmitter_TransmitServer,
) error {
	for {
		sessionID := uuid.New().String()
		mean := rand.Float64()*20 - 10
		std := rand.Float64()*1.2 + 0.3

		log.Printf("New connection - SessionID: %s, Mean: %f, STD: %f", sessionID, mean, std)
		for {
			frequency := rand.NormFloat64()*std + mean
			timestamp := timestamppb.Now()
			resp := &pb.Response{
				SessionId: sessionID,
				Frequency: frequency,
				Timestamp: timestamp,
			}
			if err := stream.Send(resp); err != nil {
				return err
			}
			time.Sleep(time.Minute)

		}
	}

	return nil
}
