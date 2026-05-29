package database

import (
	"fmt"
	"jonathangunawan30/expense-tracker/config"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQL(cfg *config.DatabaseConfig, log *logrus.Logger) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error getting sql.DB from gorm: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db
}
