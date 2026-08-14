package config

import (
	"time"
)

type GRPCConfig struct {
	Port            string        `envconfig:"PORT" required:"true"`
	Timeout         time.Duration `envconfig:"TIMEOUT" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}
