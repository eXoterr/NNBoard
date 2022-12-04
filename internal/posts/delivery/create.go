package delivery

import (
	"github.com/eXoterr/NNBoard/internal/posts/models"
	"github.com/gofiber/fiber/v2"
)

func (r *HttpHandler) Create(c *fiber.Ctx) error {

	model := &models.Post{}

	err := c.BodyParser(model)

	model.Author = 0

	if err != nil {
		return err
	}

	err = r.uc.CreatePost(model)
	if err != nil {
		return c.JSON(
			fiber.Map{"error": err.Error()},
		)
	}
	return err
}
