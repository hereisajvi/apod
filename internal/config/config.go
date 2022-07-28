package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const prefix = "apod"

// Config contains the config values for the application.
type Config struct {
	APIKey              string `required:"true" split_words:"true"`
	ServerHost          string `default:"" split_words:"true"`
	ServerPort          string `default:"8080" split_words:"true"`
	PostgresUsername    string `default:"postgres" split_words:"true"`
	PostgresPassword    string `required:"true" split_words:"true"`
	PostgresHost        string `default:"postgres" split_words:"true"`
	PostgresPort        string `default:"5432" split_words:"true"`
	PostgresDatabase    string `default:"postgres" split_words:"true"`
	PostgresSSLMode     string `default:"disable" split_words:"true"`
	MigrationsDirectory string `default:"file://migrations" split_words:"true"`
}

// New parses config from environment variables and returns an instance of config or an error.
func New() (*Config, error) {
	var cfg Config

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not read config from environment variables")
	}

	return &cfg, nil
}
