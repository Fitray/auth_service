package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LoggerConfig   LoggerConfig
	GRPCConfig     GRPCConfig
	ProjectConfig  ProjectConfig
	PostgresConfig PostgresConfig
	RedisConfig    RedisConfig
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("LOGGER", &config.LoggerConfig); err != nil {
		return Config{}, err
	}

	if err := envconfig.Process("PROJECT", &config.ProjectConfig); err != nil {
		return Config{}, err
	}

	if err := envconfig.Process("GRPC", &config.GRPCConfig); err != nil {
		return Config{}, err
	}

	if err := envconfig.Process("POSTGRES", &config.PostgresConfig); err != nil {
		return Config{}, err
	}

	if err := envconfig.Process("REDIS", &config.RedisConfig); err != nil {
		return Config{}, err
	}

	return config, nil
}
