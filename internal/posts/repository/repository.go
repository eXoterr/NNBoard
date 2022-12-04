package repository

import "github.com/eXoterr/NNBoard/internal/posts/models"

type Repository interface {
	Create(title string, body string, author uint) error
	GetByTitle(title string) (*models.Post, error)
}
