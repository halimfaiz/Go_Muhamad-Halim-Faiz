package database

import (
	"Praktikum/config"
	"Praktikum/models"
)

var (
	users []models.User
	user  models.User
)

func GetUsers() (interface{}, error) {
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	if e := config.DB.Find(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	if e := config.DB.Save(user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUserById(id int) (interface{}, error) {
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUserById(id int, updatedUser *models.User) (*models.User, error) {
	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	if e := config.DB.Save(&user).Error; e != nil {
		return nil, e
	}
	return &user, nil
}
