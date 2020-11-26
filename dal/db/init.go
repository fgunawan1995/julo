package db

import (
	"github.com/jmoiron/sqlx"
)

type DAL struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *DAL {
	return &DAL{
		DB: db,
	}
}
