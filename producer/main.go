package main

import (
	"flag"
	"time"

	"github.com/viswals/core/infrastructure/rabbitmq"
	"github.com/viswals/core/pkg/logger"
	"github.com/viswals/producer/config"
	"github.com/viswals/producer/usecase"
	"go.uber.org/zap"
)

func main() {
	// load app configuration from .env file
	config, err := config.LoadConfig()
	if err != nil {
		panic("failed to load configuration")
	}

	// initialize logger
	logger, err := logger.NewDefaultLogger()
	if err != nil {
		panic("failed to initialize logger")
	}
	logger.Info("Logger initialized")

	// filepath flag
	filepathFlag := flag.String("filepath", "./users.csv", "Path to the CSV file")
	flag.Parse()

	if filepathFlag == nil {
		logger.Fatal("no csv file specified")
	}

	filepath := *filepathFlag

	if filepath == "" {
		logger.Error("csv file path can not be empty")
	}

	// initialize rabbitmq connection
	rmq, err := rabbitmq.New(config.RabbitMQURL, 10, time.Second*5)
	if err != nil {
		logger.Fatal("failed to initialize RabbitMQ", zap.Error(err))
	}
	defer rmq.Channel.Close()
	defer rmq.Close()

	// declare queue
	queue, err := rmq.QueueDeclare(config.QueueName, rabbitmq.WithDurable(true))
	if err != nil {
		logger.Fatal("failed to declare queue", zap.Error(err))
	}
	logger.Info("Queue declared", zap.Any("queue", queue))

	// initialize service layer
	usecase := usecase.New(
		rmq,
		usecase.WithLogger(logger),
	)

	// publish CSV data to the queue
	err = usecase.PublishCSVDataToQueue(filepath, queue.Name)
	if err != nil {
		logger.Fatal("failed to publish CSV rows to queue", zap.Error(err))
	}

	logger.Info("All CSV rows have been successfully published to the queue")
}
