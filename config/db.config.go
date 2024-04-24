package config

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectUsersDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", os.Getenv("DB"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
