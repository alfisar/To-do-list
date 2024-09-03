package controller

import (
	"todolist/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}
func (obj *Controller) Simple(ctx *fiber.Ctx) error {

	ctx.Status(fasthttp.StatusOK).JSON(domain.ErrorData{
		Status:  "success",
		Code:    0,
		Message: "Welcome to API To Do List version 1.0, enjoy and chersss :)",
	})
	return nil
}
