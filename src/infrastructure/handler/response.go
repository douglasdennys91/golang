package handler

import (
	"github.com/labstack/echo/v4"
)

type handlerJSON struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Status bool `json:"status"`
}

func RenderJSON(ctx echo.Context, code int, body interface{}, status bool) error {
	return ctx.JSON(code, handlerJSON{Code: code, Data: body, Status: status})
}