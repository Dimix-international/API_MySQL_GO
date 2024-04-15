package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

var (
	cfg  Config
	once sync.Once
)

type Config struct {
	HTTPServer ServerHTTP
	Database   DB
}

type ServerHTTP struct {
	Port        string        `env:"HTTP_SERVER_PORT"`
	Address     string        `env:"HTTP_SERVER_ADDRESS"`
	Timeout     time.Duration `env:"HTTP_SERVER_TIMEOUT"`
	IdleTimeout time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT"`
}

type DB struct {
	User     string `env:"DB_USER,required,notEmpty"`
	Name     string `env:"DB_NAME,required,notEmpty"`
	Net      string `env:"DB_NET,required,notEmpty"`
	Password string `env:"DB_PASSWORD,required,notEmpty"`
	Addr     string `env:"DB_ADDRESS,required,notEmpty"`
}

func MustLoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	once.Do(func() {
		if err := env.Parse(&cfg); err != nil {
			fmt.Printf("%+v\n", err)
		}
	})

	return cfg
}
