package service

import (
	"log/slog"

	"github.com/Dimix-international/API_MySQL_GO/db"
)

type UserService struct {
	log         *slog.Logger
	userStorage db.StorageUser
}

func NewUserService(log *slog.Logger, userStorage db.StorageUser) *UserService {
	return &UserService{log: log, userStorage: userStorage}
}
