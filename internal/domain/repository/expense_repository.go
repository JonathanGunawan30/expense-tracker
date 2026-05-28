package repository

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/model"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Save(ctx context.Context, db *gorm.DB, expense *model.Expense) (*model.Expense, error)
	GetAll(ctx context.Context, db *gorm.DB, userID int) ([]*model.Expense, error)
	GetDetail(ctx context.Context, db *gorm.DB, expenseID, userID int) (*model.Expense, error)
	Update(ctx context.Context, db *gorm.DB, expense *model.Expense) (*model.Expense, error)
	Delete(ctx context.Context, db *gorm.DB, expenseID, userID int) error
}
