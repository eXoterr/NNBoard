package auth

import (
	"github.com/eXoterr/NNBoard/internal/auth/delivery"
	"github.com/eXoterr/NNBoard/internal/auth/repository"
	"github.com/eXoterr/NNBoard/internal/auth/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Setup(db *sqlx.DB, logger logrus.Logger) delivery.HttpDelivery {

	repo := repository.NewSqliteRepository(db, &logger)
	uc := usecase.NewAuthUC(repo)
	ah := delivery.NewAuthHandler(uc)
	return ah
}
