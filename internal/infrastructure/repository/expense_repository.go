package repository

import (
	"context"
	"errors"
	errDomain "jonathangunawan30/expense-tracker/internal/domain/errors"
	"jonathangunawan30/expense-tracker/internal/domain/model"
	"jonathangunawan30/expense-tracker/internal/domain/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ExpenseRepository struct {
	log *logrus.Logger
}

func NewExpenseRepository(log *logrus.Logger) repository.ExpenseRepository {
	return &ExpenseRepository{log: log}
}

func (e *ExpenseRepository) Save(ctx context.Context, db *gorm.DB, expense *model.Expense) (*model.Expense, error) {
	if err := db.WithContext(ctx).Preload("Category").Create(&expense).Error; err != nil {
		return nil, err
	}

	expenseDetail, err := e.GetDetail(ctx, db, expense.ID, expense.UserID)
	if err != nil {
		return nil, err
	}

	return expenseDetail, nil
}

func (e *ExpenseRepository) GetAll(ctx context.Context, db *gorm.DB, userID int) ([]*model.Expense, error) {
	var expenses []*model.Expense
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Preload("Category").Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (e *ExpenseRepository) GetDetail(ctx context.Context, db *gorm.DB, expenseID, userID int) (*model.Expense, error) {
	var expense model.Expense
	if err := db.WithContext(ctx).Where("id = ? AND user_id = ?", expenseID, userID).Preload("Category").First(&expense).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errDomain.ErrExpenseNotFound
		}
		return nil, err
	}
	return &expense, nil
}

func (e *ExpenseRepository) Update(ctx context.Context, db *gorm.DB, expense *model.Expense) (*model.Expense, error) {
	result := db.WithContext(ctx).Where("id = ? AND user_id = ?", expense.ID, expense.UserID).Updates(&expense)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errDomain.ErrExpenseNotFound
	}

	expenseDetail, err := e.GetDetail(ctx, db, expense.ID, expense.UserID)
	if err != nil {
		return nil, err
	}

	return expenseDetail, nil
}

func (e *ExpenseRepository) Delete(ctx context.Context, db *gorm.DB, expenseID, userID int) error {
	result := db.WithContext(ctx).Where("id = ? AND user_id = ?", expenseID, userID).Delete(&model.Expense{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errDomain.ErrExpenseNotFound
	}

	return nil
}
