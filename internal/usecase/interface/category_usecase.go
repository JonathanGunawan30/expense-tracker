package _interface

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/entity/request"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
)

type CategoryUsecase interface {
	Create(ctx context.Context, request *request.CategoryCreateRequest) (*response.CategoryResponse, error)
	GetAll(ctx context.Context, userID int) ([]*response.CategoryResponse, error)
	GetByID(ctx context.Context, categoryID, userID int) (*response.CategoryResponse, error)
	Update(ctx context.Context, request *request.CategoryUpdateRequest) (*response.CategoryResponse, error)
	Delete(ctx context.Context, categoryID, userID int) error
}
