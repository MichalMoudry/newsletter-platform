package config

import (
	"newsletter-platform/config/errors"
	"os"

	"github.com/spf13/viper"
)

type Environment string

// Structure that defines the app config.
type Config struct {
	Port int
	Environment
	ConnectionString string
}

const (
	Dev  Environment = "dev"
	Test Environment = "test"
	Prod Environment = "prod"
)

// This function reads app's configuration from a config file.
func ReadConfigFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if error := viper.ReadInConfig(); error != nil {
		return Config{}, error
	}

	env := Environment(os.Getenv("APP_ENVIRONMENT"))
	if env == "" {
		env = Environment(viper.GetString("Environment"))
	}

	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" && env == Dev {
		connectionString = viper.GetString("ConnectionString")
	} else if connectionString == "" && env != Dev {
		return Config{}, errors.ErrMissingConnectionString
	}

	return Config{
		Port:             viper.GetInt("Port"),
		Environment:      env,
		ConnectionString: connectionString,
	}, nil
}
