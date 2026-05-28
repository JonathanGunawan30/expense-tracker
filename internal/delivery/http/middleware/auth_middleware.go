package middleware

import (
	"jonathangunawan30/expense-tracker/internal/delivery/http/helper"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AuthMiddleware(log *logrus.Logger, config *viper.Viper, next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		apiKey := r.Header.Get("X-API-Key")
		secretKey := config.GetString("X_API_KEY")

		if apiKey != secretKey {
			log.Warnf("Unauthorized access attempt with API Key: %s", apiKey)
			status := http.StatusUnauthorized
			helper.WriteJSON(log, w, status, response.WebResponseError{
				Code:   status,
				Status: http.StatusText(status),
				Error:  "Unauthorized",
			})
			return
		}

		next(w, r, params)
	}
}
