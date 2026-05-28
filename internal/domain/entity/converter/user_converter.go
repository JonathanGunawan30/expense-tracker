package converter

import (
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	"jonathangunawan30/expense-tracker/internal/domain/model"
)

func ToUserResponse(user *model.User) *response.UserResponse {
	return &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []*model.User) []*response.UserResponse {
	var responses []*response.UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(user))
	}
	return responses
}
