package main

import (
	"github.com/Dimix-international/API_MySQL_GO/internal/config"
	"github.com/Dimix-international/API_MySQL_GO/internal/logger"
	"github.com/Dimix-international/API_MySQL_GO/internal/server"
)

func main() {
	cfg := config.MustLoadConfig()
	log := logger.SetupLogger(cfg.Env)
	server.NewAPIServer(cfg, nil, log).Run()
}
