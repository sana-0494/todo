package store

import (
	"fmt"
	"todo/configs"

	"github.com/jmoiron/sqlx"
)

type PgStore struct {
	db *sqlx.DB
}

func NewPostgresStore(cfg configs.Config) (*PgStore, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PgStore{
		db: db,
	}, nil
}
