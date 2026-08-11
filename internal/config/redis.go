package config

import "time"

type RedisConfig struct {
	Host     string        `envconfig:"HOST" default:"localhost"`
	Port     string        `envconfig:"PORT" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	DB       int           `envconfig:"DB" default:"0"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
	PoolSize int           `envconfig:"POOL_SIZE" default:"10"`
	Username string        `envconfig:"USERNAME" default:"app"`
}
