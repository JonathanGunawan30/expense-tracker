package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name      string `mapstructure:"APP_NAME"`
	Port      int    `mapstructure:"APP_PORT"`
	SecretKey string `mapstructure:"X_API_KEY"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

func NewConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigFile(".env")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		logrus.Warnf("Config file .env not found, using environment variables: %v", err)
	}

	return config
}
