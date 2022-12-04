package usecase

import "github.com/eXoterr/NNBoard/internal/posts/models"

type UseCase interface {
	CreatePost(post *models.Post) error
	// Register(model models.Post) error
}
