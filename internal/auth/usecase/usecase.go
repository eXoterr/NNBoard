package usecase

import "github.com/eXoterr/NNBoard/internal/auth/models"

type UseCase interface {
	LoginIn(login string, password string) error
	Register(model models.User) error
}
