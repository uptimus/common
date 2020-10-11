package etcd

import (
	"time"

	"go.uber.org/zap"

	"github.com/uptimus/common/config"
	"github.com/uptimus/common/logging"
)

type Config struct {
	config.PrintableInfoConfig

	DialTimeout time.Duration
	Endpoints   []string
}

func (config *Config) PrintInfo() {
	logging.Info("EtcdConfig",
		zap.Strings("Endpoints", config.Endpoints),
		zap.Duration("DialTimeout", config.DialTimeout),
	)
}
