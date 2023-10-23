package database

import (
	"database/sql"
	"fmt"
	"go-server/internal/config"

	_ "github.com/lib/pq"
)

// database client
type DB struct {
	Client *sql.DB
}

// Connect - connect to database
func Connect(cfg *config.Config) (*DB, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbUrl, cfg.DbPort, cfg.DbUsername, cfg.DbPassword, cfg.DbName,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{Client: db}, nil
}
