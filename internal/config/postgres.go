package config

import "time"

type PostgresConfig struct {
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DB       string `envconfig:"DB" required:"true"`
	Port     string `envconfig:"PORT" required:"true"`
	Host     string `envconfig:"HOST" default:"localhost"`

	Timeout time.Duration `envconfig:"TIMEOUT" required:"true"`

	MaxConns        int32         `envconfig:"MAX_CONNS" default:"10"`
	MinConns        int32         `envconfig:"MIN_CONNS" default:"2"`
	MaxConnLifetime time.Duration `envconfig:"MAX_CONN_LIFETIME" default:"1h"`
	MaxConnIdleTime time.Duration `envconfig:"MAX_CONN_IDLE_TIME" default:"30m"`
}
