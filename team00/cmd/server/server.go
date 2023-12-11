package main

import (
	"fmt"
	"go_team00/api/pb"
	connector_db "go_team00/internal/connector-db"
	loggerCreate "go_team00/internal/logger-create"
	"go_team00/internal/service"
	"go_team00/models"
	"net"

	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Options struct {
	Host   string `short:"h" long:"host" description:"the server host" default:"localhost" env:"HOST"`
	Port   int64  `short:"p" long:"port" description:"the server port" default:"8789" env:"PORT"`
	Log    string `long:"logger-create" description:"logger-create output" default:"debug" env:"LOG"`
	DbHost string `long:"dbhost" description:"the db server host" default:"localhost" env:"DB_HOST"`
	DbPort string `long:"dbport" description:"the db server port" default:"5432" env:"DB_PORT"`
	KTime  int    `long:"k-time" description:"time factor" default:"100" env:"K_TIME"`
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		return
	}
	logger := loggerCreate.InitLogger(opts.Log)
	db := connector_db.New(logger)

	err = db.AutoMigrate(&models.Session{})
	if err != nil {
		logger.Fatal("Error auto migrate", zap.Error(err))
	}

	transmitterServer := service.NewTransmitterServer(logger, opts.KTime, db)
	grpcServer := grpc.NewServer()
	pb.RegisterTransmitterServer(grpcServer, transmitterServer)

	address := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("cannot start server: ", zap.Error(err))
	}

	logger.Info("start server", zap.String("host", opts.Host), zap.Int64("port", opts.Port))

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal("cannot start server: ", zap.Error(err))
	}
}
