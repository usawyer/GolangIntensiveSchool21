package service

import (
	"go_team00/api/pb"
	"go_team00/models"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type TransmitterServer struct {
	pb.UnimplementedTransmitterServer
	logger *zap.Logger
	kTime  int
	db     *gorm.DB
}

func NewTransmitterServer(logger *zap.Logger, kTime int, db *gorm.DB) *TransmitterServer {
	return &TransmitterServer{logger: logger, kTime: kTime, db: db}
}

func (server *TransmitterServer) Transmit(
	req *pb.TransmitterRequest,
	stream pb.Transmitter_TransmitServer,
) error {
	sessionID := uuid.New().String()
	mean := rand.Float64()*20 - 10  //[-10; 10]
	std := rand.Float64()*1.2 + 0.3 //[0.3; 1.5]
	server.logger.Info(
		"New connection",
		zap.String("SessionID", sessionID),
		zap.Float64("Mean", mean),
		zap.Float64("std", std))
	var session models.Session
	server.db.FirstOrCreate(&session, models.Session{
		SessionId: sessionID,
		ServMean:  mean,
		ServStd:   std,
		Status:    "research",
	})
	for {
		frequency := rand.NormFloat64()*std + mean
		timestamp := timestamppb.Now()
		resp := &pb.TransmitterResponse{
			SessionId: sessionID,
			Frequency: frequency,
			Timestamp: timestamp,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * time.Duration(server.kTime))

	}
	return nil
}
