package request

type CategoryUpdateRequest struct {
	ID     int    `json:"id" validate:"required"`
	UserID int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required,max=255"`
}
