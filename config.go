package validator

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Password     string `envconfig:"MONGODB_PASSWORD"    json:"-"`
	AwsAuthToken string `envconfig:"AWS_AUTH_TOKEN"    json:"-"`
	Host         string `envconfig:"MONGODB_BIND_ADDR"    json:"-"`
}

var cfg *Config

// Get - configures the application and returns the cfg
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		Password:     "",
		AwsAuthToken: "",
		Host:         "db-postgres-nyc1-1111-do-user-111111-0.db.ondigitalocean.com:password=admin",
	}

	return cfg, envconfig.Process("", cfg)
}
