package router

import "github.com/gofiber/fiber/v2"

func simpleRouter(v1 fiber.Router) {
	controll := SimpleInit()
	v1.Get("", controll.Simple)
}
