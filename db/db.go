package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB
}

func NewDb(cnf mysql.Config) (*DB, error) {
	db, err := sql.Open("mysql", cnf.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
