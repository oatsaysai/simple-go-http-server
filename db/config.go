package db

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost          string
	DBPort          string
	DBUsername      string
	DBPassword      string
	DBName          string
	MaxOpenConns    int32
	MaxConnLifetime int32
}

func InitConfig() (*Config, error) {
	config := &Config{
		DBHost:          viper.GetString("postgresql.DBHost"),
		DBPort:          viper.GetString("postgresql.DBPort"),
		DBUsername:      viper.GetString("postgresql.DBUsername"),
		DBPassword:      viper.GetString("postgresql.DBPassword"),
		DBName:          viper.GetString("postgresql.DBName"),
		MaxOpenConns:    viper.GetInt32("postgresql.MaxOpenConns"),
		MaxConnLifetime: viper.GetInt32("postgresql.MaxConnLifetime"),
	}
	if config.DBHost == "" {
		config.DBHost = "localhost"
	}
	if config.DBPort == "" {
		config.DBPort = "5432"
	}
	if config.DBUsername == "" {
		config.DBUsername = "postgres"
	}
	if config.DBPassword == "" {
		config.DBPassword = "postgres"
	}
	if config.DBName == "" {
		config.DBName = "simple_db"
	}
	return config, nil
}
