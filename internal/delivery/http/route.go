package http

import (
	"jonathangunawan30/expense-tracker/internal/delivery/http/handler"
	"jonathangunawan30/expense-tracker/internal/delivery/http/middleware"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewRouter(userHandler *handler.UserHandler, categoryHandler *handler.CategoryHandler, expenseHandler *handler.ExpenseHandler, log *logrus.Logger, config *viper.Viper) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", middleware.AuthMiddleware(log, config, userHandler.Register))
	router.GET("/api/users", middleware.AuthMiddleware(log, config, userHandler.GetAll))

	router.POST("/api/categories", middleware.AuthMiddleware(log, config, categoryHandler.Create))
	router.GET("/api/categories", middleware.AuthMiddleware(log, config, categoryHandler.GetAll))
	router.GET("/api/categories/:categoryID", middleware.AuthMiddleware(log, config, categoryHandler.GetDetail))
	router.PUT("/api/categories/:categoryID", middleware.AuthMiddleware(log, config, categoryHandler.Update))
	router.DELETE("/api/categories/:categoryID", middleware.AuthMiddleware(log, config, categoryHandler.Delete))

	router.POST("/api/expenses", middleware.AuthMiddleware(log, config, expenseHandler.Create))
	router.GET("/api/expenses", middleware.AuthMiddleware(log, config, expenseHandler.GetAll))
	router.GET("/api/expenses/:expenseID", middleware.AuthMiddleware(log, config, expenseHandler.GetDetail))
	router.PUT("/api/expenses/:expenseID", middleware.AuthMiddleware(log, config, expenseHandler.Update))
	router.DELETE("/api/expenses/:expenseID", middleware.AuthMiddleware(log, config, expenseHandler.Delete))

	return router
}