package app

import "github.com/spf13/viper"

type Config struct {
	MaxStringLength int64
}

func InitConfig() (*Config, error) {
	config := &Config{
		MaxStringLength: viper.GetInt64("App.MaxStringLength"),
	}
	if config.MaxStringLength == 0 {
		config.MaxStringLength = 10
	}
	return config, nil
}
