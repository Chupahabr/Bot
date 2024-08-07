package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	skinsTable    = "skins"
	usersTable    = "users"
	stickersTable = "stickers"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSL      string
}

func ConnectDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSL))

	if err != nil {
		return nil, err
	}

	return db, nil
}
