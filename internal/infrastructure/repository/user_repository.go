package repository

import (
	"context"
	errDomain "jonathangunawan30/expense-tracker/internal/domain/errors"
	"jonathangunawan30/expense-tracker/internal/domain/model"
	"jonathangunawan30/expense-tracker/internal/domain/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) repository.UserRepository {
	return &UserRepository{log: log}
}

func (u *UserRepository) Save(ctx context.Context, db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		if errDomain.IsDuplicateEntry(err) {
			return nil, errDomain.ErrEmailDuplicate
		}
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetAll(ctx context.Context, db *gorm.DB) ([]*model.User, error) {
	var users []*model.User
	if err := db.WithContext(ctx).Select("id", "name", "email", "created_at", "updated_at").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
