package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User
var mapUsers map[int]User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	UserId := GetIdFromParam(c)

	if val, ok := mapUsers[UserId]; ok {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success get user by id",
			"user":     val,
		})
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "error get user by id",
		"user":     nil,
	})
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	UserId := GetIdFromParam(c)
	for index, val := range users {
		if val.Id == UserId {
			users = append(users[:index], users[index+1:]...)
			delete(mapUsers, UserId)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "succes delete user by id ",
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message":          "Error delete user by id",
		"errorDescription": "id not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	UserId := GetIdFromParam(c)
	updatedUser := User{}
	c.Bind(&updatedUser)
	updatedUser.Id = UserId

	for index, val := range users {
		if val.Id == UserId {
			users[index] = updatedUser
			mapUsers[UserId] = updatedUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user by id",
			})
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"messages":         "error update user by id",
		"errorDescription": "Not Found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func GetIdFromParam(c echo.Context) int {
	id, _ := strconv.Atoi(c.Param("id"))
	return id
}

// ---------------------------------------------------
func main() {
	mapUsers = make(map[int]User)
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
