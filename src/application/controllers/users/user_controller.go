package users

import (
	"delivery-app/src/application/services"
	"delivery-app/src/domain/repositories"
	response2 "delivery-app/src/infrastructure/handler/response"
	"delivery-app/src/infrastructure/utils/json"
	"github.com/labstack/echo/v4"
)

var (
	UserCTRL IUserController = &UserController{}
)

type UserController struct {
	repositories.UserRepository
	services.UserService
}

type IUserController interface {
	GetAllUsers(ctx echo.Context) error
	CreatedUser(ctx echo.Context) error
	GetUserByID(ctx echo.Context) error
	GetUserParam(ctx echo.Context) error
	UpdateUserByID(ctx echo.Context) error
	RemoveUserByID(ctx echo.Context) error
}

type Param struct {
	Email string `json:"email"`
}

func (ctrl *UserController) GetAllUsers(ctx echo.Context) error {
	response, err := ctrl.GetUsers()
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	return response2.RenderJSON(ctx, 200, response, true)
}

func (ctrl *UserController) CreatedUser(ctx echo.Context) error {
	data, err := json.ParserByte(ctx.Request().Body)
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	body, err := ctrl.Saved(data)
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	return response2.RenderJSON(ctx, 201, body, true)
}

func (ctrl *UserController) GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")
	response, err := ctrl.GetUser(id)
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	return response2.RenderJSON(ctx, 200, response, true)
}

func (ctrl *UserController) GetUserParam(ctx echo.Context) error {
	param, err := json.ParserJSON(ctx.Request().Body)
	response, err := ctrl.GetUserByParam(param)
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	return response2.RenderJSON(ctx, 200, response, true)
}

func (ctrl *UserController) UpdateUserByID(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := json.ParserByte(ctx.Request().Body)
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	response, err := ctrl.Updated(id, data)
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	return response2.RenderJSON(ctx, 200, response, true)
}

func (ctrl *UserController) RemoveUserByID(ctx echo.Context) error {
	id := ctx.Param("id")
	response, err := ctrl.DeleteUser(id)
	if err != nil {
		return response2.RenderJSON(ctx, 500, err.Error(), false)
	}
	return response2.RenderJSON(ctx, 200, response, true)
}
