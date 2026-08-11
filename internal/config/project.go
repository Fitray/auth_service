package config

type ProjectConfig struct {
	ProjectRootPath string `envconfig:"ROOT" required:"true"`
}
