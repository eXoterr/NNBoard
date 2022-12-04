package delivery

import (
	"github.com/eXoterr/NNBoard/internal/auth/models"
	"github.com/gofiber/fiber/v2"
)

func (l *AuthHandler) Login(c *fiber.Ctx) error {

	model := &models.User{}

	err := c.BodyParser(model)

	if err != nil {
		return err
	}
	err = l.uc.LoginIn(model.Login, model.Password)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "success"})
}
