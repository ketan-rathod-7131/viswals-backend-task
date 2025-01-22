package usecase

import (
	"github.com/viswals/core/infrastructure/rabbitmq"
	"github.com/viswals/core/interfaces"
	"github.com/viswals/core/pkg/logger"
)

type ProducerUsecase struct {
	rmq    interfaces.IQueueService
	logger interfaces.ILogger
}

// set default values for producer
func (p *ProducerUsecase) setDefaults() {
	if p.logger == nil {
		logger, err := logger.NewDefaultLogger()
		if err != nil {
			panic(err)
		}

		p.logger = logger
	}
}

type Option func(*ProducerUsecase)

func WithRabbitMQ(rmq *rabbitmq.RabbitMQ) Option {
	return func(p *ProducerUsecase) {
		p.rmq = rmq
	}
}

func WithLogger(logger interfaces.ILogger) Option {
	return func(p *ProducerUsecase) {
		p.logger = logger
	}
}

func New(rmq interfaces.IQueueService, options ...Option) *ProducerUsecase {
	usecase := &ProducerUsecase{
		rmq: rmq,
	}

	for _, option := range options {
		option(usecase)
	}

	usecase.setDefaults()

	return usecase
}
