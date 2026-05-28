package repository

import (
	"context"
	"errors"
	domainErr "jonathangunawan30/expense-tracker/internal/domain/errors"
	"jonathangunawan30/expense-tracker/internal/domain/model"
	"jonathangunawan30/expense-tracker/internal/domain/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) repository.CategoryRepository {
	return &CategoryRepository{log: log}
}

func (c *CategoryRepository) Save(ctx context.Context, db *gorm.DB, category *model.Category) (*model.Category, error) {
	if err := db.WithContext(ctx).Create(category).Error; err != nil {
		if domainErr.IsFKViolation(err) {
			return nil, domainErr.ErrUserNotFound
		}
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepository) GetDetail(ctx context.Context, db *gorm.DB, categoryID, userID int) (*model.Category, error) {
	var category model.Category
	if err := db.WithContext(ctx).Where("id = ? AND user_id = ?", categoryID, userID).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErr.ErrCategoryOrUserNotFound
		}
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepository) GetAll(ctx context.Context, db *gorm.DB, userID int) ([]*model.Category, error) {
	var categories []*model.Category
	if err := db.WithContext(ctx).Select("id", "user_id", "name", "created_at", "updated_at").Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryRepository) Update(ctx context.Context, db *gorm.DB, category *model.Category) (*model.Category, error) {
	result := db.WithContext(ctx).Where("id = ? AND user_id = ?", category.ID, category.UserID).Updates(category)

	if result.Error != nil {
		if domainErr.IsFKViolation(result.Error) {
			return nil, domainErr.ErrUserNotFound
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domainErr.ErrCategoryOrUserNotFound
	}

	return category, nil
}

func (c *CategoryRepository) Delete(ctx context.Context, db *gorm.DB, categoryID, userID int) error {
	result := db.WithContext(ctx).Where("id = ? AND user_id = ?", categoryID, userID).Delete(&model.Category{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainErr.ErrCategoryOrUserNotFound
	}

	return nil
}
