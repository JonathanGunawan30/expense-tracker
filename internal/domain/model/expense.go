package model

import "time"

type Expense struct {
	ID          int        `gorm:"column:id;primaryKey;<-:false;autoIncrement"`
	UserID      int        `gorm:"column:user_id"`
	CategoryID  int        `gorm:"column:category_id"`
	Amount      int64      `gorm:"column:amount;not null"`
	Title       string     `gorm:"column:title;not null"`
	Description *string    `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`

	User     User     `gorm:"foreignKey:UserID;references:ID"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}

func (e *Expense) TableName() string {
	return "expenses"
}
