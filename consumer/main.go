package main

import (
	"context"
	"flag"
	"net/http"
	"sync"
	"time"

	"github.com/viswals/consumer/config"
	controller "github.com/viswals/consumer/controller/http"
	"github.com/viswals/consumer/usecase"
	"github.com/viswals/core/infrastructure/encryption"
	"github.com/viswals/core/infrastructure/postgres"
	"github.com/viswals/core/infrastructure/rabbitmq"
	"github.com/viswals/core/infrastructure/redis"
	"github.com/viswals/core/interfaces"
	"github.com/viswals/core/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// load app configuration
	config, err := config.LoadConfig()
	if err != nil {
		panic("failed to load configuration")
	}

	// initialize Logger
	logger, err := logger.NewDefaultLogger()
	if err != nil {
		panic("failed to initialize logger")
	}
	logger.Info("Logger initialized")

	// take consume flag
	var startConsumerService bool
	consumeFlag := flag.Bool("consume", false, "consume used to specify wheather to consume data from rabbitmq queue")
	flag.Parse()
	if consumeFlag != nil {
		startConsumerService = *consumeFlag
	}

	// initialize RabbitMQ
	rmq, err := rabbitmq.New(config.RabbitMQURL, 10, time.Second*5)
	if err != nil {
		logger.Fatal("failed to initialize RabbitMQ", zap.Error(err))
	}
	defer rmq.Channel.Close()
	defer rmq.Close()

	queue, err := rmq.QueueDeclare(config.QueueName, rabbitmq.WithDurable(true))
	if err != nil {
		logger.Fatal("failed to declare queue", zap.Error(err))
	}
	logger.Info("Queue declared", zap.Any("queue", queue))

	postgresDB, err := postgres.NewPostgres(config.DBConfig)
	if err != nil {
		logger.Fatal("failed to initialize postgres", zap.Error(err))
	}

	var cm interfaces.ICacheService
	cm, err = redis.New(config.RedisConfig)
	if err != nil {
		logger.Error("failed to initialize redis", zap.Error(err))
		cm = redis.NewNoOpCache()
	}

	em, err := encryption.New([]byte(config.EncryptionKey))
	if err != nil {
		logger.Fatal("failed to initialize encryption manager", zap.Error(err))
	}

	usecase := usecase.New(postgresDB, rmq, em, usecase.WithLogger(logger), usecase.WithCacheManager(cm))

	httpMux := http.NewServeMux()
	httpController := controller.New(usecase, controller.WithHttpMux(httpMux), controller.WithHttpPort(config.HttpPort))

	// wg for grpc and http servers
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		if startConsumerService {
			logger.Info("starting consumer service")
			usecase.ConsumeUserData(context.Background(), config.QueueName)
		} else {
			logger.Info("consume flag not passed, skipping consumer service")
		}
	}()

	go func() {
		defer wg.Done()
		logger.Info("starting http server", zap.String("port", config.HttpPort))
		if err := httpController.Start(); err != nil {
			logger.Error("cannot run http server", zap.Error(err), zap.String("port", config.HttpPort))
			panic(err)
		}
	}()

	wg.Wait()
	logger.Info("consumer stopped!")
	// Initialize redis cache layer for consumer
}
