package status

import (
	"delivery-app/src/infrastructure/handler/response"
	"github.com/labstack/echo/v4"
)

func RouterStatus(ctx echo.Context) error {
	return response.RenderJSON(ctx, 200, "Echo Server started successfully!", true)
}