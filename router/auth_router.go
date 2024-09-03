package router

import (
	"todolist/application/authentification/controller"
	"todolist/helper"
	"todolist/internal/handler"
	"todolist/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type authRouter struct {
	Controller controller.AuthControllerContract
}

func NewAuthRouter(authController controller.AuthControllerContract) *authRouter {
	return &authRouter{
		Controller: authController,
	}
}

func (obj *authRouter) authRouters(v1 fiber.Router) {
	v1.Post("/registration", middleware.Validation(handler.HandlerRegistration, helper.ValidationDataUser), obj.Controller.Registration)
	v1.Post("/login", middleware.Validation(handler.HandlerLogin, helper.ValidationLogin), obj.Controller.Login)

}
