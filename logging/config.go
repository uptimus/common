package logging

import (
	"fmt"

	"gitlab.wvservices.com/waves/gateways/gw-commons/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelKey  = "level"
	FormatKey = "format"
)

type Config struct {
	config.PrintableInfoConfig

	// "debug", "info", "warn", "error", "dpanic", "panic", and "fatal"
	Level string
	// "console" and json
	Format string
	// "dispatcher", "provider", etc
	Name string
	// map[string]interface{}{"foo": 1, "bar": 2}
	Payloads   map[string]interface{}
	StackTrace bool
}

func (config *Config) PrintInfo() {
	Info("LoggingConfig",
		zap.String(LevelKey, config.Level),
		zap.String(FormatKey, config.Format),
	)
}

func (config *Config) ToZapConfig() (*zap.Config, error) {
	var level = zap.NewAtomicLevel()
	if err := level.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, err
	}

	switch config.Format {
	case "console":
		break
	case "json":
		break
	// Set default logging format
	case "":
		config.Format = "json"
	default:
		return nil, fmt.Errorf("logging format should 'console' or 'json', but received: %s", config.Format)
	}

	return &zap.Config{
		Level:             level,
		DisableStacktrace: !config.StackTrace,
		Development:       false,
		Encoding:          config.Format,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		InitialFields: config.Payloads,
	}, nil
}
