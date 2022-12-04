package repository

import (
	"github.com/eXoterr/NNBoard/internal/posts/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func (r *PostRepository) Create(title string, body string, author uint) error {
	_, err := r.DB.Exec("INSERT INTO Posts (title,body,author) VALUES($1,$2,$3);", title, body, author)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetByTitle(title string) (*models.Post, error) {
	model := &models.Post{}
	err := r.DB.Get(model, "SELECT * FROM `Users` WHERE title = $1 LIMIT 1", title)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func NewSqliteRepository(db *sqlx.DB, logger *logrus.Logger) Repository {
	return &PostRepository{
		DB:     db,
		Logger: logger,
	}
}
