package model

import "time"

type User struct {
	ID        int        `gorm:"column:id;primaryKey;<-:false;autoIncrement"`
	Email     string     `gorm:"column:email;unique;not null"`
	Name      string     `gorm:"column:name;not null"`
	Password  string     `gorm:"column:password;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`

	Categories []Category `gorm:"foreignKey:UserID"`
	Expenses   []Expense  `gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "users"
}
