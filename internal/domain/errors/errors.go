package errors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailDuplicate         = errors.New("email already exists")
	ErrExpenseNotFound        = errors.New("expense not found")
	ErrCategoryOrUserNotFound = errors.New("category or user not found")
)

func IsFKViolation(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1452
	}
	return false
}

func IsDuplicateEntry(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062
	}
	return false
}
