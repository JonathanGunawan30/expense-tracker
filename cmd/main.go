package main

import (
	"jonathangunawan30/expense-tracker/config"
	"jonathangunawan30/expense-tracker/internal/infrastructure/database"
)

func main() {
	viperConfig := config.NewConfig()
	log := config.NewLogger(viperConfig)
	dbConfig := config.NewDatabase(viperConfig)
	db := database.NewMySQL(dbConfig, log)
	validate := config.NewValidator()

	config.Bootstrap(viperConfig, log, db, validate)
}
