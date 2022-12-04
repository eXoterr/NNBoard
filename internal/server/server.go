package server

import (
	"github.com/eXoterr/NNBoard/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Server struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
	Config *config.Config
	Router *fiber.App
}

func (s *Server) Run() error {

	srv := fiber.New(fiber.Config{
		ReadTimeout:  s.Config.ReadTimeout,
		WriteTimeout: s.Config.WriteTimeout,
	})

	s.Router = srv

	db, err := SetupDatabase()

	if err != nil {
		return err
	}

	s.DB = db

	err = s.SetupHandlers(s.Logger, s.DB)

	if err != nil {
		return err
	}

	return s.Router.Listen(s.Config.ListenAddr)
}
