package database

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (*model.User, error)
	GetAllUsers() (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user model.User) (*model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetAllUsers() (*model.User, error) {
	var user model.User
	err := u.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
