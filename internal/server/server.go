package server

import (
	"fmt"
	"go-server/internal/config"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type server struct {
	config *config.Config
}

func Start() {

	log.SetLevel(log.DebugLevel)
	log.Info("setting up go-server...")

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error getting config:", err.Error())
	}

	// TODO - connect to a database

	s := &server{
		config: cfg,
	}

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:     s.routes(),
		IdleTimeout: time.Minute,
	}

	log.Infof("running server on port %d...", cfg.ServerPort)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("error starting go-server:", err.Error())
	}
}
