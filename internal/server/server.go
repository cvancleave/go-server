package server

import (
	"fmt"
	"go-server/internal/config"
	"go-server/internal/database"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type server struct {
	config   *config.Config
	database *database.DB
}

func Start() {

	log.SetLevel(log.DebugLevel)
	log.Info("setting up go-server...")

	// get server config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error getting config: %s", err.Error())
	}

	// connect to a database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}
	defer db.Client.Close()

	s := &server{
		config:   cfg,
		database: db,
	}

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:     s.routes(),
		IdleTimeout: time.Minute,
	}

	log.Infof("running server on port %d...", cfg.ServerPort)

	// start server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error starting go-server: %s", err.Error())
	}
}
