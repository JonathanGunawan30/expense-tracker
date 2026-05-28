package request

type ExpenseCreateRequest struct {
	UserID      int     `json:"user_id" validate:"required"`
	CategoryID  int     `json:"category_id" validate:"required"`
	Amount      int64   `json:"amount" validate:"required"`
	Title       string  `json:"title" validate:"required,max=255"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}
