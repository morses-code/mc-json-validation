package validator

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	Password     string `envconfig:"MONGODB_PASSWORD"    json:"-"`
	AwsAuthToken string `envconfig:"AWS_AUTH_TOKEN"    json:"-"`
}

const VIPER_CONFIG = ".env"

var cfg *Config

// Get - configures the application and returns the cfg
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	viper.SetConfigFile(VIPER_CONFIG)
	viper.ReadInConfig()

	cfg = &Config{
		Password:     "#^n@JnmcyzwS91B%$!d2Wb#CVnZt8D3L",
		AwsAuthToken: viper.Get("AWS_AUTH_TOKEN").(string),
	}

	return cfg, envconfig.Process("", cfg)
}
