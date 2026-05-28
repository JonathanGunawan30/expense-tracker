package helper

import (
	"errors"
	errDomain "jonathangunawan30/expense-tracker/internal/domain/errors"
	"net/http"
)

func DomainErrorToHTTPStatus(err error) int {
	switch {
	case errors.Is(err, errDomain.ErrEmailDuplicate):
		return http.StatusConflict
	case errors.Is(err, errDomain.ErrUserNotFound):
		return http.StatusNotFound
	case errors.Is(err, errDomain.ErrExpenseNotFound):
		return http.StatusNotFound
	case errors.Is(err, errDomain.ErrCategoryOrUserNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
