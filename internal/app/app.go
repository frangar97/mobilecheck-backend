package app

import (
	"log"
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/config"
	"github.com/frangar97/mobilecheck-backend/internal/handler"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/frangar97/mobilecheck-backend/internal/service"
	"github.com/frangar97/mobilecheck-backend/pkg/postgres"
)

func Run() {
	cfg, err := config.InitConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewClient(cfg.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)

	handlers := handler.NewHandler(services)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handlers.Init(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
