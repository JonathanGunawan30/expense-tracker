package response

import "time"

type ExpenseResponse struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	CategoryID  int        `json:"category_id"`
	Category    *Category  `json:"category"`
	Amount      int64      `json:"amount"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
