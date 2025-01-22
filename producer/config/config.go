package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitMQURL string
	QueueName   string
	LoggerLevel string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	// load environment variables from .env file, if available
	_ = godotenv.Load()

	return &Config{
		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		QueueName:   getEnv("RABBITMQ_QUEUE", "example-queue"),
		LoggerLevel: getEnv("LOGGER_LEVEL", "debug"),
	}, nil
}

// getEnv retrieves an environment variable or a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
