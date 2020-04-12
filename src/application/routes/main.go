package routes

import (
	"delivery-app/src/application/routes/status"
	"delivery-app/src/application/routes/users"
	"github.com/labstack/echo/v4"
)

func Router(router *echo.Echo) {
	router.GET("/", status.RouterStatus)
	users.Routes(router)
}
