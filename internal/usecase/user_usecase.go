package usecase

import (
	"context"
	"jonathangunawan30/expense-tracker/internal/domain/entity/converter"
	"jonathangunawan30/expense-tracker/internal/domain/entity/request"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	"jonathangunawan30/expense-tracker/internal/domain/model"
	"jonathangunawan30/expense-tracker/internal/domain/repository"
	_interface "jonathangunawan30/expense-tracker/internal/usecase/interface"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase struct {
	db             *gorm.DB
	log            *logrus.Logger
	userRepository repository.UserRepository
}

func NewUserUsecase(db *gorm.DB, log *logrus.Logger, userRepository repository.UserRepository) _interface.UserUsecase {
	return &UserUsecase{db: db, log: log, userRepository: userRepository}
}

func (u *UserUsecase) Register(ctx context.Context, request *request.UserRegisterRequest) (*response.UserResponse, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorf("failed to generate hash password: %v", err)
		return nil, err
	}

	user := &model.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: string(passwordHash),
	}

	save, err := u.userRepository.Save(ctx, tx, user)
	if err != nil {
		u.log.Errorf("failed to save user: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		u.log.Errorf("failed to commit transaction: %v", err)
		return nil, err
	}

	return converter.ToUserResponse(save), nil
}

func (u *UserUsecase) GetAll(ctx context.Context) ([]*response.UserResponse, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	users, err := u.userRepository.GetAll(ctx, tx)
	if err != nil {
		u.log.Errorf("failed to get all users: %v", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		u.log.Errorf("failed to commit transaction: %v", err)
		return nil, err
	}

	return converter.ToUserResponses(users), nil
}
