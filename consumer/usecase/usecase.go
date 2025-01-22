package usecase

import (
	"context"

	"github.com/viswals/consumer/usecase/repository/database"
	"github.com/viswals/core/infrastructure/postgres"
	"github.com/viswals/core/infrastructure/rabbitmq"
	"github.com/viswals/core/infrastructure/redis"
	"github.com/viswals/core/interfaces"
	"github.com/viswals/core/models"
	"github.com/viswals/core/pkg/logger"
	"github.com/viswals/core/pkg/utils"
)

type IConsumerRepository interface {
	CreateUser(ctx context.Context, user models.User) (id string, err error)
	GetUserByEmail(ctx context.Context, email string) (user models.User, err error)
	GetUserById(ctx context.Context, id int64) (user models.User, err error)
	GetAllUsers(ctx context.Context, pagination utils.PaginationParams, filters []utils.Filter) (users []models.User, totalUsers int, err error)
}

type ConsumerUsecase struct {
	rmq    interfaces.IQueueService
	logger interfaces.ILogger
	db     IConsumerRepository
	em     interfaces.IEncryptionService
	cm     interfaces.ICacheService
}

func (p *ConsumerUsecase) setDefaults() {
	if p.logger == nil {
		logger, err := logger.NewDefaultLogger()
		if err != nil {
			panic(err)
		}

		p.logger = logger
	}

	// if cache is not defined by default, then use the default no operation cache implementation
	if p.cm == nil {
		p.cm = redis.NewNoOpCache()
	}
}

type Option func(*ConsumerUsecase)

func WithRabbitMQ(rmq *rabbitmq.RabbitMQ) Option {
	return func(p *ConsumerUsecase) {
		p.rmq = rmq
	}
}

func WithCacheManager(em interfaces.ICacheService) Option {
	return func(p *ConsumerUsecase) {
		p.cm = em
	}
}

func WithLogger(logger interfaces.ILogger) Option {
	return func(p *ConsumerUsecase) {
		p.logger = logger
	}
}

func New(db postgres.Postgres, rmq *rabbitmq.RabbitMQ, em interfaces.IEncryptionService, options ...Option) *ConsumerUsecase {

	usecase := &ConsumerUsecase{
		em:  em,
		rmq: rmq,
	}

	for _, option := range options {
		option(usecase)
	}

	usecase.setDefaults()

	// initialize repo layer
	consumerDB := database.New(db.DB, usecase.logger)
	usecase.db = consumerDB

	return usecase
}
