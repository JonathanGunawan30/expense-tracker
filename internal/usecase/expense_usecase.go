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

type ExpenseUsecase struct {
	db                 *gorm.DB
	log                *logrus.Logger
	categoryRepository repository.CategoryRepository
	expenseRepository  repository.ExpenseRepository
}

func NewExpenseUsecase(db *gorm.DB, log *logrus.Logger, categoryRepository repository.CategoryRepository, expenseRepository repository.ExpenseRepository) _interface.ExpenseUsecase {
	return &ExpenseUsecase{db: db, log: log, categoryRepository: categoryRepository, expenseRepository: expenseRepository}
}

func (e *ExpenseUsecase) Create(ctx context.Context, request request.ExpenseCreateRequest) (*response.ExpenseResponse, error) {
	tx := e.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	_, err := e.categoryRepository.GetDetail(ctx, tx, request.CategoryID, request.UserID)
	if err != nil {
		return nil, err
	}

	expense := &model.Expense{
		UserID:      request.UserID,
		CategoryID:  request.CategoryID,
		Amount:      request.Amount,
		Title:       request.Title,
		Description: request.Description,
	}

	save, err := e.expenseRepository.Save(ctx, tx, expense)
	if err != nil {
		e.log.Errorf("failed to create expense: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		e.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToExpenseResponse(save), nil
}

func (e *ExpenseUsecase) GetAll(ctx context.Context, userID int) ([]*response.ExpenseResponse, error) {
	tx := e.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	expenses, err := e.expenseRepository.GetAll(ctx, tx, userID)
	if err != nil {
		e.log.Errorf("failed to get expenses: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		e.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToExpenseResponses(expenses), nil
}

func (e *ExpenseUsecase) GetDetail(ctx context.Context, expenseID, userID int) (*response.ExpenseResponse, error) {
	tx := e.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	expense, err := e.expenseRepository.GetDetail(ctx, tx, expenseID, userID)
	if err != nil {
		e.log.Errorf("failed to get expense detail: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		e.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToExpenseResponse(expense), nil
}

func (e *ExpenseUsecase) Update(ctx context.Context, request request.ExpenseUpdateRequest) (*response.ExpenseResponse, error) {
	tx := e.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	_, err := e.categoryRepository.GetDetail(ctx, tx, request.CategoryID, request.UserID)
	if err != nil {
		return nil, err
	}

	expense := &model.Expense{
		ID:          request.ID,
		UserID:      request.UserID,
		CategoryID:  request.CategoryID,
		Amount:      request.Amount,
		Title:       request.Title,
		Description: request.Description,
	}

	update, err := e.expenseRepository.Update(ctx, tx, expense)
	if err != nil {
		e.log.Errorf("failed to update expense: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		e.log.Errorf("failed to commit database transaction: %v", err)
		return nil, err
	}

	return converter.ToExpenseResponse(update), nil
}

func (e *ExpenseUsecase) Delete(ctx context.Context, expenseID, userID int) error {
	tx := e.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := e.expenseRepository.Delete(ctx, tx, expenseID, userID); err != nil {
		e.log.Errorf("failed to delete expense: %v", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		e.log.Errorf("failed to commit database transaction: %v", err)
		return err
	}

	return nil
}
