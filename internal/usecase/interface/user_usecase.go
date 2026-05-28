package _interface

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/entity/request"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
)

type UserUsecase interface {
	Register(ctx context.Context, request *request.UserRegisterRequest) (*response.UserResponse, error)
	GetAll(ctx context.Context) ([]*response.UserResponse, error)
}
