package delivery

import (
	"github.com/eXoterr/NNBoard/internal/auth/usecase"
	"github.com/gofiber/fiber/v2"
)

type HttpDelivery interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type AuthHandler struct {
	uc usecase.UseCase
}

func NewAuthHandler(uc usecase.UseCase) HttpDelivery {
	return &AuthHandler{
		uc: uc,
	}
}
