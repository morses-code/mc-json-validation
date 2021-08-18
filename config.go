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
		Password:     "g!h3gMsWFJ5psGRJu#EF2KzyVDZ9y^aq",
		AwsAuthToken: "realtoken",
	}

	return cfg, envconfig.Process("", cfg)
}
