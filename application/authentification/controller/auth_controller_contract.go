package controller

import "github.com/gofiber/fiber/v2"

type AuthControllerContract interface {
	Registration(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
