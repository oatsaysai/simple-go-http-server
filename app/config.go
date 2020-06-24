package app

import "github.com/spf13/viper"

type Config struct {
	MaxStringLength int
}

func InitConfig() (*Config, error) {
	config := &Config{
		MaxStringLength: viper.GetInt("App.MaxStringLength"),
	}
	if config.MaxStringLength == 0 {
		config.MaxStringLength = 10
	}
	return config, nil
}
