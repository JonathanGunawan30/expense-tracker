package config

import (
	"fmt"
	"jonathangunawan30/expense-tracker/internal/delivery/http"
	"jonathangunawan30/expense-tracker/internal/delivery/http/handler"
	"jonathangunawan30/expense-tracker/internal/infrastructure/repository"
	"jonathangunawan30/expense-tracker/internal/usecase"
	netHttp "net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Bootstrap(config *viper.Viper, log *logrus.Logger, db *gorm.DB, validate *validator.Validate) {
	userRepository := repository.NewUserRepository(log)
	categoryRepository := repository.NewCategoryRepository(log)
	expenseRepository := repository.NewExpenseRepository(log)

	userUsecase := usecase.NewUserUsecase(db, log, userRepository)
	categoryUsecase := usecase.NewCategoryUsecase(db, log, categoryRepository)
	expenseUsecase := usecase.NewExpenseUsecase(db, log, categoryRepository, expenseRepository)

	userHandler := handler.NewUserHandler(log, validate, userUsecase)
	categoryHandler := handler.NewCategoryHandler(log, validate, categoryUsecase)
	expenseHandler := handler.NewExpenseHandler(log, validate, expenseUsecase)

	router := http.NewRouter(userHandler, categoryHandler, expenseHandler, log, config)

	port := config.GetInt("APP_PORT")
	if port == 0 {
		port = 3000
	}

	address := fmt.Sprintf(":%d", port)
	log.Infof("Server started at %s", address)
	err := netHttp.ListenAndServe(address, router)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
