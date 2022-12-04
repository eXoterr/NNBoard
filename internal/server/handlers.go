package server

import (
	"github.com/eXoterr/NNBoard/internal/auth"
	"github.com/eXoterr/NNBoard/internal/posts"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func (s *Server) SetupHandlers(logger *logrus.Logger, db *sqlx.DB) error {

	// Setup repository
	// authRepo := auth.NewSqliteRepository(db, logger)
	// postRepo := posts.NewSqliteRepository(db, logger)
	//
	// //Setup UseCase
	// authUC := auth.NewAuthUC(authRepo)
	// postsUC := posts.NewPostsUC(postRepo)
	//
	// //setupHandler
	// authHandler := http.NewAuthHandler(authUC)
	//

	authHandler := auth.Setup(db, *logger)
	s.Router.Post("/register", authHandler.Register)
	s.Router.Post("/login", authHandler.Login)

	postsHandler := posts.Setup(db, *logger)
	s.Router.Post("/posts/create", postsHandler.Create)

	return nil
}
