package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(config *viper.Viper) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	level, err := logrus.ParseLevel(config.GetString("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)

	return logger
}
