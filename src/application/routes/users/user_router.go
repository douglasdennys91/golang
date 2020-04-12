package users

import (
	"delivery-app/src/application/controllers/users"
	"github.com/labstack/echo/v4"
)

func Routes(router *echo.Echo) {
	router.GET("/users", users.UserCTRL.GetAllUsers)
	router.POST("/users", users.UserCTRL.CreatedUser)
	router.GET("/users/:id", users.UserCTRL.GetUserByID)
	router.POST("/users/param", users.UserCTRL.GetUserParam)
	router.PUT("/users/:id", users.UserCTRL.UpdateUserByID)
	router.DELETE("/users/:id", users.UserCTRL.RemoveUserByID)
}