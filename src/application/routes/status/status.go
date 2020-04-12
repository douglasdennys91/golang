package status

import (
	"delivery-app/src/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func RouterStatus(ctx echo.Context) error {
	return handler.RenderJSON(ctx, 200, "Echo Server started successfully!", true)
}