package config

type LoggerConfig struct {
	Level  string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
	Format string `envconfig:"FORMAT" required:"true"`
}
