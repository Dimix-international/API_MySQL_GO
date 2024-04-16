package db

import "database/sql"

type UserStorageStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorageStorage {
	return &UserStorageStorage{db: db}
}
