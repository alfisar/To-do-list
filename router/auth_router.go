package router

import (
	"todolist/application/authentification/controller"

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
	v1.Post("/registration", obj.Controller.Registration)
}
