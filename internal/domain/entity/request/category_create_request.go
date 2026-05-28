package request

type CategoryCreateRequest struct {
	UserID int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required,max=255"`
}
