package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(UserUsecase usecase.UserUsecase) *userController {
	return &userController{userUsecase: UserUsecase}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userUsecase.GetAllUsers()
	if err != nil {
		return c.JSON(400, "Failed to Get All Users")
	}
	return c.JSON(200, map[string]interface{}{
		"data": users,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, err.Error())
	}

	user, err := u.userUsecase.CreateUser(*user)
	if err != nil {
		return c.JSON(400, "Failed To Create User")
	}

	return c.JSON(200, map[string]interface{}{
		"data": user,
	})
}
