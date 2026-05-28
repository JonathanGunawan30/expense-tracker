package converter

import (
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	"jonathangunawan30/expense-tracker/internal/domain/model"
)

func ToExpenseResponse(expense *model.Expense) *response.ExpenseResponse {
	return &response.ExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		CategoryID:  expense.CategoryID,
		Category:    ToCategoryDive(&expense.Category),
		Amount:      expense.Amount,
		Title:       expense.Title,
		Description: expense.Description,
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
	}
}

func ToCategoryDive(category *model.Category) *response.Category {
	return &response.Category{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToExpenseResponses(expenses []*model.Expense) []*response.ExpenseResponse {
	var responses []*response.ExpenseResponse
	for _, value := range expenses {
		responses = append(responses, ToExpenseResponse(value))
	}
	return responses
}
