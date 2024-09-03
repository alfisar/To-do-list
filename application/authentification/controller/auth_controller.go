package controller

import (
	"log"
	"todolist/application/authentification/service"
	"todolist/domain"
	"todolist/internal/consts"
	"todolist/internal/errorhandler"
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
	request := domain.User{}
	errData := ctx.BodyParser(&request)
	if errData != nil {
		log.Printf("Error parsing data request on controller registration : %s", errData.Error())
		err := errorhandler.ErrValidation(errData)
		response.WriteResponse(ctx, response.Response{}, err, err.Code)
		return nil
	}

	user, err := c.serv.Registration(ctx.Context(), ctx.Request(), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.Code)
		return nil
	}

	user.Password = ""
	resp := response.ResponseSuccess(user, consts.SuccessRegister)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
