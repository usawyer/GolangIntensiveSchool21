package main

import (
	"context"
	"fmt"
	"go_team00/api/pb"
	connector_db "go_team00/internal/connector-db"
	logger_create "go_team00/internal/logger-create"
	"go_team00/internal/service"
	"go_team00/models"
	"log"
	"math"

	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

type Options struct {
	Host            string  `short:"h" long:"host" description:"the server host" default:"localhost" env:"HOST"`
	Port            int64   `short:"p" long:"port" description:"the server port" default:"8789" env:"PORT"`
	Log             string  `long:"logger-create" description:"logger-create output" default:"debug" env:"LOG"`
	DbHost          string  `description:"the db server host" default:"localhost" env:"DB_HOST"`
	DbPort          string  `description:"the db server port" default:"5432" env:"DB_PORT"`
	AnomalyK        float64 `short:"k" description:"the STD anomaly coefficient" default:"2.0" env:"K_ANOMALY"`
	PredictDataSize int     `long:"size" short:"s" description:"he data set size for prediction processing" default:"100" env:"SIZE_PRED"`
}

var (
	opts Options
)

type clientData struct {
	logger            *zap.Logger
	transmitterClient pb.TransmitterClient
	clientDb          *gorm.DB
	opts              Options
}

func (c *clientData) ProcessMessages() {
	stream, err := c.transmitterClient.Transmit(context.Background(), &pb.TransmitterRequest{})
	if err != nil {
		c.logger.Fatal("Transmit failed", zap.Error(err))
	}

	stats := service.NewClientStats()

	for {
		resp, err := stream.Recv()
		if err != nil {
			c.logger.Fatal("Error receiving message", zap.Error(err))
		}
		if stats.GetState() == service.StateTypeCollection {
			if !stats.UpdateData(resp.Frequency, c.opts.PredictDataSize) {
				var session models.Session
				if err := c.clientDb.Where("session_id = ?", resp.SessionId).Find(&session).Error; err != nil {
					c.logger.Fatal("failed to fetch sessions: ", zap.Error(err))
				}
				c.clientDb.Model(&session).Updates(models.Session{
					ClientStd:  stats.GetSTD(),
					ClientMean: stats.GetMean(),
					Status:     "anomaly detected",
					ClientK:    c.opts.AnomalyK,
				})
			}
			stats.LogInfo(c.logger)
		}
		if stats.GetState() == service.StateTypeDetection {
			if isAnomaly(resp.Frequency, stats.GetMean(), stats.GetSTD(), c.opts.AnomalyK) {
				var session models.Session
				c.clientDb.FirstOrCreate(&session, models.Session{SessionId: resp.SessionId})

				c.clientDb.Create(&models.Anomaly{
					SessionID: session.ID,
					Frequency: resp.Frequency,
					Timestamp: resp.Timestamp.AsTime(),
				})

				c.logger.Info("Anomaly detected",
					zap.Float64("Frequency", resp.Frequency),
					zap.String("SessionID", resp.SessionId),
					zap.Time("time", resp.Timestamp.AsTime()),
				)
			}
		}
	}
}

func isAnomaly(frequency, mean, std, k float64) bool {
	deviation := math.Abs(frequency - mean)
	return deviation > k*std
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Println(err.Error())
	}
	logger := logger_create.InitLogger(opts.Log)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error(err.Error())
		}
	}(logger)

	servAddress := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	conn, err := grpc.Dial(servAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot dial server", zap.Error(err))
	}
	defer conn.Close()

	transmitterClient := pb.NewTransmitterClient(conn)
	logger.Info("connect to server", zap.String("server", servAddress))
	db := connector_db.New(logger)

	data := clientData{
		logger:            logger,
		transmitterClient: transmitterClient,
		clientDb:          db,
		opts:              opts,
	}

	err = db.AutoMigrate(&models.Anomaly{})

	if err != nil {
		logger.Fatal("Error auto migrate", zap.Error(err))
	}

	data.ProcessMessages()
}
