package app

import (
	log "github.com/oatsaysai/simple-go-http-server/log"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.SetDefault("LogLevel", "debug")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func createLoggerForTest() (log.Logger, error) {
	initConfig()

	logLevel := viper.GetString("Log.Level")
	logLevel = log.NormalizeLogLevel(logLevel)

	logColor := viper.GetBool("Log.Color")
	logJSON := viper.GetBool("Log.JSON")

	logger, err := log.NewLogger(&log.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logLevel,
		ConsoleJSONFormat: logJSON,
		Color:             logColor,
	}, log.InstanceZapLogger)
	if err != nil {
		return nil, err
	}
	return logger, nil
}
