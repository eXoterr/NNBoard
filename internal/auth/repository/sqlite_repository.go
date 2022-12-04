package repository

import (
	"github.com/eXoterr/NNBoard/internal/auth/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func (r *AuthRepository) Create(login string, password string) error {
	_, err := r.DB.Exec("INSERT INTO Users (login,password) VALUES($1,$2);", login, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetByLogin(login string) (*models.User, error) {
	model := &models.User{}
	err := r.DB.Get(model, "SELECT * FROM `Users` WHERE login = $1 LIMIT 1", login)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func NewSqliteRepository(db *sqlx.DB, logger *logrus.Logger) Repository {
	return &AuthRepository{
		DB:     db,
		Logger: logger,
	}
}
