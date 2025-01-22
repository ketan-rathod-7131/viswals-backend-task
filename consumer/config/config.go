package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/viswals/core/infrastructure/postgres"
	"github.com/viswals/core/infrastructure/redis"
)

type Config struct {
	RabbitMQURL   string
	QueueName     string
	LoggerLevel   string
	EncryptionKey string
	DBConfig      *postgres.DbConfig
	RedisConfig   *redis.RedisConfig
	HttpPort      string
}

// LoadConfig loads the app configuration from the .env file.
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	dbConfig := &postgres.DbConfig{
		MigrationsPath:   getEnv("MIGRATION_PATH", "./core/documents/migrations"),
		ConnectionString: getEnv("POSTGRES_CONNECTION_URL", ""),
		MaxConnRetries:   10,
		MaxOpenConns:     20,
	}

	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	redisPort, _ := strconv.Atoi(getEnv("REDIS_PORT", "6379"))

	redisConfig := &redis.RedisConfig{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Password: getEnv("REDIS_PASSWORD", ""),
		Port:     redisPort,
		DB:       redisDB,
	}

	return &Config{
		RabbitMQURL:   getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		QueueName:     getEnv("RABBITMQ_QUEUE", "example-queue"),
		EncryptionKey: getEnv("ENCRYPTION_KEY", ""),
		LoggerLevel:   getEnv("LOGGER_LEVEL", "debug"),
		HttpPort:      getEnv("HTTP_PORT", "8080"),
		DBConfig:      dbConfig,
		RedisConfig:   redisConfig,
	}, nil
}

// getEnv retrieves an environment variable or a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
