package _interface

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/entity/request"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
)

type ExpenseUsecase interface {
	Create(ctx context.Context, request request.ExpenseCreateRequest) (*response.ExpenseResponse, error)
	GetAll(ctx context.Context, userID int) ([]*response.ExpenseResponse, error)
	GetDetail(ctx context.Context, expenseID, userID int) (*response.ExpenseResponse, error)
	Update(ctx context.Context, request request.ExpenseUpdateRequest) (*response.ExpenseResponse, error)
	Delete(ctx context.Context, expenseID, userID int) error
}
