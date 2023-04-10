package routes

import (
	"Praktikum/constants"
	"Praktikum/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)
	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.CreateUserController)

	// Route / to handler function/ User
	user := e.Group("/user", middleware.JWT([]byte(constants.SECRET_JWT)))
	user.GET("s", controllers.GetUsersController)
	user.GET("s/:id", controllers.GetUserController)
	user.DELETE("s/:id", controllers.DeleteUserController)
	user.PUT("s/:id", controllers.UpdateUserController)

	//Book
	book := e.Group("/books", middleware.JWT([]byte(constants.SECRET_JWT)))
	book.GET("/:id", controllers.GetBookController)
	book.GET("", controllers.GetBooksController)
	book.POST("", controllers.CreateBookController)
	book.DELETE("/:id", controllers.DeleteBookController)
	book.PUT("/:id", controllers.UpdateBookController)

	//Blog
	blog := e.Group("/blogs", middleware.JWT([]byte(constants.SECRET_JWT)))
	blog.GET("/:id", controllers.GetBlogController)
	blog.GET("", controllers.GetBlogsController)
	blog.POST("", controllers.CreateBlogController)
	blog.DELETE("/:id", controllers.DeleteBlogController)
	blog.PUT("/:id", controllers.UpdateBlogController)
	return e
}
