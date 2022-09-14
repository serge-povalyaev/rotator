package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"bannerRotator/internal/config"
	"bannerRotator/internal/logger"
	"bannerRotator/internal/rabbit"
	"bannerRotator/internal/repository"
	"bannerRotator/internal/server"
	"bannerRotator/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config/rotator.yaml", "Путь до файла конфигурации")
}

func main() {
	flag.Parse()
	conf := config.ReadConfig(configPath)
	log := logger.New(conf.Logger.Level)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	connection, err := sqlx.Connect("postgres", conf.DB.CreateDSN())
	if err != nil {
		cancel()
		log.Fatal(err.Error())
	}

	producer := rabbit.NewProducer(log, conf.Rabbit)
	err = producer.Connect()
	if err != nil {
		cancel()
		log.Fatal(err.Error())
	}

	rotatorService := service.NewRotatorService(
		log,
		repository.NewBannerRepository(connection),
		repository.NewBannerToSlotRepository(connection),
		repository.NewSlotRepository(connection),
		repository.NewSocialGroupRepository(connection),
		repository.NewStatRepository(connection),
		repository.NewTotalStatRepository(connection),
		producer,
	)

	grpcServer := server.NewGRPCServer(rotatorService, log)

	go func() {
		<-ctx.Done()

		grpcServer.Stop()
		if err = connection.Close(); err != nil {
			cancel()
			log.Fatal(err.Error())
		}
		if err = producer.Disconnect(); err != nil {
			cancel()
			log.Fatal(err.Error())
		}
	}()

	if err := grpcServer.Start(conf.GRPC); err != nil {
		cancel()
		log.Fatal(err.Error())
	}

	defer cancel()
}
