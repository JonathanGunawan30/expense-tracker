package converter

import (
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	"jonathangunawan30/expense-tracker/internal/domain/model"
)

func ToCategoryResponse(category *model.Category) *response.CategoryResponse {
	return &response.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		UserID:    category.UserID,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []*model.Category) []*response.CategoryResponse {
	var responses []*response.CategoryResponse
	for _, value := range categories {
		responses = append(responses, ToCategoryResponse(value))
	}
	return responses
}
