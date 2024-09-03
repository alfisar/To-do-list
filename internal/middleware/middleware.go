package middleware

import (
	"log"
	"todolist/internal/errorhandler"
	"todolist/internal/response"

	"github.com/gofiber/fiber/v2"
)

func Validation[T any](parse func(c *fiber.Ctx) (T, error), validate func(T) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request, err := parse(c)
		if err != nil {
			log.Printf("Error parsing data on  middleware : %s", err.Error())
			err := errorhandler.ErrValidation(err)
			response.WriteResponse(c, response.Response{}, err, fiber.StatusBadRequest)
			return nil
		}

		err = validate(request)
		if err != nil {
			log.Printf("Error validation data on middleware : %s", err.Error())
			err := errorhandler.ErrValidation(err)
			response.WriteResponse(c, response.Response{}, err, fiber.StatusBadRequest)
			return nil
		}

		c.Locals("validatedData", request)
		return c.Next()
	}
}
