package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new logger.
// Available modes: "production" and "development".
// Available levels: error, warn, info, debug.
func NewLogger(mode, level string) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "message",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// parse string level representation to zapcore.Level, default is InfoLevel
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// Config for constructing logger
	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Add options for "development" mode
	switch mode {
	case "development":
		loggerConfig.EncoderConfig.CallerKey = "caller"
		loggerConfig.EncoderConfig.StacktraceKey = "stacktrace"
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		loggerConfig.Development = true
		loggerConfig.Encoding = "console"
	default:
		loggerConfig.Development = false
		loggerConfig.Encoding = "json"
	}

	logger, _ := loggerConfig.Build()
	return logger
}
