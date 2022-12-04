package repository

import "github.com/eXoterr/NNBoard/internal/auth/models"

type Repository interface {
	Create(login string, pwdhash string) error
	GetByLogin(login string) (*models.User, error)
}
