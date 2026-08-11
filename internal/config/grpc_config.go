package config

import (
	"time"
)

type GRPCConfig struct {
	Port    string        `envconfig:"PORT" required:"true"`
	Timeout time.Duration `envconfig:"TIMEOUT" required:"true"`
}
