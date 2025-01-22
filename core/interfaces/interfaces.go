package interfaces

import (
	"context"
	"time"

	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/viswals/core/infrastructure/rabbitmq"
)

type ILogger interface {
	Debug(message string, args ...zap.Field)
	Info(message string, args ...zap.Field)
	Warn(message string, args ...zap.Field)
	Error(message string, args ...zap.Field)
	Fatal(message string, args ...zap.Field)
}

// ICacheService interface defines the methods for caching
type ICacheService interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

// IEncryptionService is the interface for encryption and hashing.
type IEncryptionService interface {
	Encrypt(data string) (string, error)
	Decrypt(data string) (string, error)
	Hash(data string) (string, error)
	CompareHash(data, hash string) (bool, error)
}

type IQueueService interface {
	ConsumeWithContext(ctx context.Context, queue string, options ...rabbitmq.ConsumeOption) (<-chan amqp.Delivery, error)
	PublishWithContext(ctx context.Context, options ...rabbitmq.PublishOption) error
	QueueDeclare(name string, options ...rabbitmq.QueueOption) (amqp.Queue, error)
	Qos(prefetchCount, prefetchSize int, global bool) error
}
