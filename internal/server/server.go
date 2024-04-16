package server

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dimix-international/API_MySQL_GO/internal/config"
	"github.com/Dimix-international/API_MySQL_GO/internal/models"
)

type APIServer struct {
	cfg     config.Config
	db      *sql.DB
	log     *slog.Logger
	closers []models.CloseFunc
}

func NewAPIServer(cfg config.Config, db *sql.DB, log *slog.Logger) *APIServer {
	return &APIServer{
		cfg: cfg,
		db:  db,
		log: log,
	}
}

func (s *APIServer) Run() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.launchServer(); err != nil {
			s.log.Error(fmt.Sprintf("Stop server: %v", err))
			exit <- syscall.SIGTERM
			close(exit)
		}
	}()

	<-exit
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := s.Shutdown(shutdownCtx); err != nil {
		s.log.Error(fmt.Sprintf("error closing server: %v", err))
	}
}

func (s *APIServer) launchServer() error {
	return nil
}

func (s *APIServer) AddCloser(closer models.CloseFunc) {
	s.closers = append(s.closers, closer)
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	for _, fn := range s.closers {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}
