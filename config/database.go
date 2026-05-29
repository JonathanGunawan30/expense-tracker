package config

import (
	"github.com/spf13/viper"
)

func NewDatabase(config *viper.Viper) *DatabaseConfig {
	return &DatabaseConfig{
		Host:     config.GetString("DB_HOST"),
		Port:     config.GetInt("DB_PORT"),
		User:     config.GetString("DB_USER"),
		Password: config.GetString("DB_PASSWORD"),
		Name:     config.GetString("DB_NAME"),
	}
}
