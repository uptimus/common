package config

import (
	"github.com/spf13/viper"
)

// Load config structure from filePath
func Load(cfg interface{}, filePath string) error {
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	return nil
}

type PrintableInfoConfig interface {
	PrintInfo()
}
