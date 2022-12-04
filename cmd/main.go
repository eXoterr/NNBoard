package main

import (
	"log"
	"time"

	"github.com/eXoterr/NNBoard/internal/config"
	"github.com/eXoterr/NNBoard/internal/server"
	"github.com/eXoterr/NNBoard/pkg/logger"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.NewConfig()
	err := cfg.ReadConfig()

	if err != nil {
		log.Fatalln(err)
	}

	logg, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalln(err)
	}

	cfg.WriteTimeout *= time.Second
	cfg.ReadTimeout *= time.Second

	srv := server.Server{
		DB:     &sqlx.DB{},
		Logger: logg,
		Config: cfg,
	}

	err = srv.Run()

	if err != nil {
		log.Fatalln(err)
	}

}
