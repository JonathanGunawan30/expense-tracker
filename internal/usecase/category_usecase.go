package usecase

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/entity/converter"
	"jonathangunawan30/expense-tracker/internal/domain/entity/request"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	"jonathangunawan30/expense-tracker/internal/domain/model"
	"jonathangunawan30/expense-tracker/internal/domain/repository"
	_interface "jonathangunawan30/expense-tracker/internal/usecase/interface"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryUsecase struct {
	db                 *gorm.DB
	log                *logrus.Logger
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(db *gorm.DB, log *logrus.Logger, categoryRepository repository.CategoryRepository) _interface.CategoryUsecase {
	return &CategoryUsecase{db: db, log: log, categoryRepository: categoryRepository}
}

func (c *CategoryUsecase) Create(ctx context.Context, request *request.CategoryCreateRequest) (*response.CategoryResponse, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	category := &model.Category{
		UserID: request.UserID,
		Name:   request.Name,
	}

	save, err := c.categoryRepository.Save(ctx, tx, category)
	if err != nil {
		c.log.Errorf("failed to create category: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		c.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToCategoryResponse(save), nil
}

func (c *CategoryUsecase) GetAll(ctx context.Context, userID int) ([]*response.CategoryResponse, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	categories, err := c.categoryRepository.GetAll(ctx, tx, userID)
	if err != nil {
		c.log.Errorf("failed to get all categories: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		c.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToCategoryResponses(categories), nil
}

func (c *CategoryUsecase) GetByID(ctx context.Context, categoryID, userID int) (*response.CategoryResponse, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	category, err := c.categoryRepository.GetDetail(ctx, tx, categoryID, userID)
	if err != nil {
		c.log.Errorf("failed to get category detail: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		c.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToCategoryResponse(category), nil
}

func (c *CategoryUsecase) Update(ctx context.Context, request *request.CategoryUpdateRequest) (*response.CategoryResponse, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	category := &model.Category{
		ID:     request.ID,
		UserID: request.UserID,
		Name:   request.Name,
	}

	update, err := c.categoryRepository.Update(ctx, tx, category)
	if err != nil {
		c.log.Errorf("failed to update category: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		c.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToCategoryResponse(update), nil
}

func (c *CategoryUsecase) Delete(ctx context.Context, categoryID, userID int) error {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.categoryRepository.Delete(ctx, tx, categoryID, userID)
	if err != nil {
		c.log.Errorf("failed to delete category: %v", err)
		return err
	}

	if err = tx.Commit().Error; err != nil {
		c.log.Errorf("failed to commit database transaction: %v", err)
		return err
	}

	return nil
}
