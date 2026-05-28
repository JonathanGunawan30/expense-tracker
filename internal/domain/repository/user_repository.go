package repository

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, db *gorm.DB, user *model.User) (*model.User, error)
	GetAll(ctx context.Context, db *gorm.DB) ([]*model.User, error)
}
