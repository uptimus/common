package logging

import (
	"context"

	"go.uber.org/zap"
)

const (
	LoggerKey = "logger"
)

var logger *Logger

type Logger struct {
	*zap.Logger
}

func (l *Logger) Close() error {
	return l.Sync()
}

// Init default logger
func Default() (*Logger, error) {
	if logger != nil {
		return logger, nil
	}

	l, err := Init(&Config{})
	if err != nil {
		return nil, err
	}
	logger = l
	return logger, nil
}

// Init logger from config
func Init(config *Config) (*Logger, error) {
	zapConfig, err := config.ToZapConfig()
	if err != nil {
		return nil, err
	}

	// Add caller skipper because it's wrapped
	logger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	// Set caller name (default if empty = no logger field)
	logger = logger.Named(config.Name)
	// Reset global logger
	zap.ReplaceGlobals(logger)
	return &Logger{logger}, nil
}

func (l *Logger) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, LoggerKey, l)
}

func LoggerFromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(LoggerKey).(*Logger); ok {
		return l
	} else {
		l, _ := Default()
		return l
	}
}

func (l *Logger) Named(s string) *Logger {
	return &Logger{l.Logger.Named(s)}
}

func With(fields ...zap.Field) *zap.Logger {
	return zap.L().With(fields...)
}

func Named(s string) *zap.Logger {
	return zap.L().Named(s)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return zap.L().WithOptions(opts...)
}

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zap.L().Panic(msg, fields...)
}
