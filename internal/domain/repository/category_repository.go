package repository

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(ctx context.Context, db *gorm.DB, category *model.Category) (*model.Category, error)
	GetDetail(ctx context.Context, db *gorm.DB, categoryID, userID int) (*model.Category, error)
	GetAll(ctx context.Context, db *gorm.DB, userID int) ([]*model.Category, error)
	Update(ctx context.Context, db *gorm.DB, category *model.Category) (*model.Category, error)
	Delete(ctx context.Context, db *gorm.DB, categoryID, userID int) error
}
