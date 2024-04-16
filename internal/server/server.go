package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dimix-international/API_MySQL_GO/db"
	"github.com/Dimix-international/API_MySQL_GO/internal/config"
	"github.com/Dimix-international/API_MySQL_GO/internal/handlers"
	"github.com/Dimix-international/API_MySQL_GO/internal/models"
	"github.com/Dimix-international/API_MySQL_GO/internal/service"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type APIServer struct {
	cfg         config.Config
	log         *slog.Logger
	closers     []models.CloseFunc
	router      *mux.Router
	userService *service.UserService
}

func NewAPIServer(cfg config.Config, log *slog.Logger) *APIServer {
	return &APIServer{
		cfg: cfg,
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
	s.initDBAndServices()
	s.initRoutes()

	httpServer := &http.Server{
		Handler:      s.router,
		Addr:         s.cfg.HTTPServer.Address,
		ReadTimeout:  s.cfg.HTTPServer.Timeout,
		WriteTimeout: s.cfg.HTTPServer.Timeout,
		IdleTimeout:  s.cfg.HTTPServer.IdleTimeout,
	}

	s.log.Info(fmt.Sprintf("server HTTP started no port: %v", s.cfg.HTTPServer.Port))
	s.AddCloser(httpServer.Shutdown)

	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *APIServer) initDBAndServices() error {
	client, err := db.NewDb(mysql.Config{
		User:                 s.cfg.Database.User,
		Passwd:               s.cfg.Database.Password,
		Net:                  s.cfg.Database.Net,
		DBName:               s.cfg.Database.Name,
		Addr:                 s.cfg.Database.Addr,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		return err
	}

	s.userService = service.NewUserService(s.log, db.NewUserStorage(client.DB))
	return nil
}

func (s *APIServer) initRoutes() {
	s.router = mux.NewRouter()
	subrouter := s.router.PathPrefix("/api/v1").Subrouter()

	handlers.NewUserHandler(s.log, service.NewUserService(s.log, s.userService)).RegisterUserRoutes(subrouter)
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
