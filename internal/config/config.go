package config

import (
	"fmt"

	"github.com/marisasha/warehouse-helper/internal/repository"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string            `mapstructure:"port"`
	DB      repository.Config `mapstructure:"db"`
}

func Load() (*Config, error) {

	if err := initConfig(); err != nil {
		return nil, fmt.Errorf("init config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	return viper.ReadInConfig()
}
