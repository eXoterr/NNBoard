package delivery

import (
	"github.com/eXoterr/NNBoard/internal/auth/models"
	"github.com/gofiber/fiber/v2"
)

func (r *AuthHandler) Register(c *fiber.Ctx) error {

	model := &models.User{}

	err := c.BodyParser(model)

	if err != nil {
		return err
	}

	err = r.uc.Register(*model)
	if err != nil {
		return c.JSON(
			fiber.Map{"error": err.Error()},
		)
	}
	return err
}
