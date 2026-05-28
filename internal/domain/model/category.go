package model

import "time"

type Category struct {
	ID        int        `gorm:"column:id;primaryKey;<-:false;autoIncrement"`
	UserID    int        `gorm:"column:user_id"`
	Name      string     `gorm:"column:name;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}

func (c *Category) TableName() string {
	return "categories"
}
