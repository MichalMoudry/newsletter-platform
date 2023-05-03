package config

import "github.com/spf13/viper"

type Environment string

type Config struct {
	Port int
	Environment
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

	return Config{
		Port: viper.GetInt("Port"),
	}, nil
}
