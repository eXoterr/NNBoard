package delivery

import (
	"github.com/eXoterr/NNBoard/internal/posts/usecase"
	"github.com/gofiber/fiber/v2"
)

type HttpDelivery interface {
	Create(c *fiber.Ctx) error
}

type HttpHandler struct {
	uc usecase.UseCase
}

func NewHttpHandler(uc usecase.UseCase) HttpDelivery {
	return &HttpHandler{
		uc: uc,
	}
}
