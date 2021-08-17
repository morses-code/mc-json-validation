package validator

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Password string `envconfig:"MONGODB_PASSWORD"    json:"-"`
}

var cfg *Config

// Get - configures the application and returns the cfg
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		Password: "12341234",
	}

	return cfg, envconfig.Process("", cfg)
}