package validator

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Password     string `envconfig:"MONGODB_PASSWORD"    json:"-"`
	AwsAuthToken string `envconfig:"AWS_AUTH_TOKEN"    json:"-"`
}

var cfg *Config

// Get - configures the application and returns the cfg
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		Password:     "#^n@JnmcyzwS91B%$!d2Wb#CVnZt8D3L",
		AwsAuthToken: "OH PAAANTS",
	}

	return cfg, envconfig.Process("", cfg)
}
