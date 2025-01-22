package database

import (
	"github.com/viswals/core/interfaces"
)

type ConsumerDB struct {
	DB     interfaces.ISqlDatabase
	logger interfaces.ILogger
}

func New(db interfaces.ISqlDatabase, logger interfaces.ILogger) *ConsumerDB {
	return &ConsumerDB{
		DB:     db,
		logger: logger,
	}
}
