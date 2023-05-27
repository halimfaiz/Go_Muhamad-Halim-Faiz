package routes

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/database"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := database.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
}
