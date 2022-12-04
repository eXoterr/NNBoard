package usecase

import (
	"errors"

	"github.com/eXoterr/NNBoard/internal/posts/models"
	"github.com/eXoterr/NNBoard/internal/posts/repository"
)

type PostUC struct {
	repo repository.Repository
}

func (uc *PostUC) CreatePost(model *models.Post) error {

	if isEmpty(model.Title) {
		return errors.New("empty title")
	}

	if isEmpty(model.Body) {
		return errors.New("empty body")
	}

	err := uc.repo.Create(model.Title, model.Body, model.Author)

	return err
}

func NewUC(repo repository.Repository) *PostUC {
	return &PostUC{
		repo: repo,
	}
}

func isEmpty(str string) bool {
	return len(str) == 0
}
