package usecase

import (
	"belajar-go-echo/database"
	"belajar-go-echo/model"
)

type UserUsecase interface {
	CreateUser(req model.User) (*model.User, error)
	GetAllUsers() (*model.User, error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepo database.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) CreateUser(req model.User) (*model.User, error) {
	user, err := u.userRepository.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userUsecase) GetAllUsers() (*model.User, error) {
	user, err := u.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}
