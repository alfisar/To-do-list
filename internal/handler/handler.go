package handler

import (
	"todolist/domain"

	"github.com/gofiber/fiber/v2"
)

func HandlerRegistration(c *fiber.Ctx) (domain.User, error) {
	request := domain.User{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerLogin(c *fiber.Ctx) (domain.Login, error) {
	request := domain.Login{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}
