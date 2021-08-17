package validator

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Password     string `envconfig:"MONGODB_PASSWORD"    json:"-"`
	AwsAuthToken string `envconfig:"AWS_AUTH_TOKEN"    json:"-"`
	BindAddr     string `envconfig:"MONGODB_BIND_ADDR"    json:"-"`
}

var cfg *Config

// Get - configures the application and returns the cfg
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		Password:     "",
		AwsAuthToken: "AKIAFJKR45SAWSZ5XDF3",
		BindAddr:     "localhost:27017",
	}

	return cfg, envconfig.Process("", cfg)
}
