package controller

import (
	"todolist/application/authentification/service"
	"todolist/domain"
	"todolist/internal/consts"
	"todolist/internal/response"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	serv service.AuthServiceContract
}

func NewAuthController(serv service.AuthServiceContract) *authController {
	return &authController{
		serv: serv,
	}
}

func (c *authController) Registration(ctx *fiber.Ctx) error {
	request := ctx.Locals("validatedData").(domain.User)

	user, err := c.serv.Registration(ctx.Context(), ctx.Request(), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, 400)
		return nil
	}

	user.Password = ""
	resp := response.ResponseSuccess(user, consts.SuccessRegister)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *authController) Login(ctx *fiber.Ctx) error {
	request := ctx.Locals("validatedData").(domain.Login)

	token, err := c.serv.Login(ctx.Context(), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, 400)
		return nil
	}

	resp := response.ResponseSuccessWithToken(nil, consts.SuccessLogin, token)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
