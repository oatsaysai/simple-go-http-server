package http

import "github.com/spf13/viper"

type Config struct {
	// The port to bind HTTP application API server to
	Port int

	// The number of proxies positioned in front of the API. This is used to interpret
	// X-Forwarded-For headers.
	ProxyCount int

	LogLevel string
}

func InitConfig() (*Config, error) {
	config := &Config{
		Port:     viper.GetInt("HTTPAPIServerPort"),
		LogLevel: viper.GetString("LogLevel"),
	}
	if config.Port == 0 {
		config.Port = 9092
	}
	return config, nil
}
