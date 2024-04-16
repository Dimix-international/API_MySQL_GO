package main

import (
	_ "github.com/Dimix-international/API_MySQL_GO/docs"
	"github.com/Dimix-international/API_MySQL_GO/internal/config"
	"github.com/Dimix-international/API_MySQL_GO/internal/logger"
	"github.com/Dimix-international/API_MySQL_GO/internal/server"
)

// @title Swagger User management API
// @version 1.0
// @description
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description JWT Authorization header "Bearer auth_token"
func main() {
	cfg := config.MustLoadConfig()
	log := logger.SetupLogger(cfg.Env)

	server.NewAPIServer(cfg, log).Run()
	log.Info("app finish")
}
