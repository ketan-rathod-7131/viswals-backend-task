package logger

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultZapConfig = zap.Config{
	Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
	Development: false,
	Sampling: &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	},
	Encoding: "json",
	EncoderConfig: zapcore.EncoderConfig{
		TimeKey:        "timestamp",                    // Field for timestamp
		LevelKey:       "type",                         // Field for log level (e.g., debug, info, error)
		MessageKey:     "message",                      // Field for log message
		StacktraceKey:  "",                             // Remove stack trace from log output
		CallerKey:      "caller",                       // Remove caller info
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // Lowercase level encoding
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 time format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Omit caller info
	},
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
}

// NewDefaultLogger returns you a new otelzap logger with core specific default settings
func NewDefaultLogger() (*otelzap.Logger, error) {
	zapLogger, err := defaultZapConfig.Build()
	if err != nil {
		return nil, err
	}

	logger := otelzap.New(zapLogger, otelzap.WithMinLevel(zap.DebugLevel))

	return logger, nil
}

// New creates a new otelzap logger with optional custom zap configuration.
// If no configuration is provided, it uses the default configuration.
func New(config *zap.Config) (*otelzap.Logger, error) {
	var zapLogger *zap.Logger
	var err error

	// Use custom configuration if provided, otherwise use default configuration
	if config != nil {
		zapLogger, err = config.Build()
	} else {
		zapLogger, err = defaultZapConfig.Build()
	}

	if err != nil {
		return nil, err
	}

	logger := otelzap.New(zapLogger, otelzap.WithMinLevel(zap.DebugLevel))

	return logger, nil
}
