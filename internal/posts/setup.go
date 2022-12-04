package posts

import (
	"github.com/eXoterr/NNBoard/internal/posts/delivery"
	"github.com/eXoterr/NNBoard/internal/posts/repository"
	"github.com/eXoterr/NNBoard/internal/posts/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Setup(db *sqlx.DB, logger logrus.Logger) delivery.HttpDelivery {

	repo := repository.NewSqliteRepository(db, &logger)
	uc := usecase.NewUC(repo)
	ah := delivery.NewHttpHandler(uc)
	return ah
}
