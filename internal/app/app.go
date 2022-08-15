package app

import (
	"log"
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/config"
	"github.com/frangar97/mobilecheck-backend/pkg/postgres"
)

func Run() {
	cfg, err := config.InitConfig()

	if err != nil {
		log.Fatal(err)
	}

	_, err = postgres.NewClient(cfg.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr: ":" + cfg.Port,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
