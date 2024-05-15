package db

import (
	"fmt"
	"hillel_auction/config"
	"hillel_auction/logger"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Database struct {
	log *logger.Logger
	cfg *config.Configuration
	*sqlx.DB
}

func New(log *logger.Logger, cfg *config.Configuration) *Database {
	return &Database{
		log: log,
		cfg: cfg,
	}
}

func (db *Database) Connect() error {
	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.cfg.DBUser, db.cfg.DBPassword, db.cfg.DBHost, db.cfg.DBPort, db.cfg.DBName)

	database, err := sqlx.Connect("postgres", dbConnectionString)
	if err != nil {
		db.log.Error("failed to connect to database", err)
		return err
	}

	db.DB = database

	return nil
}
