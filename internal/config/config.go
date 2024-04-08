package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	DbConfig dbConfig `yaml:"dbConfig" env-required:"true"`
	Server   Server   `yaml:"server"`
}

type dbConfig struct {
	DbUser  string `yaml:"user" env-required:"true"`
	DbPass  string `yaml:"pass" env-required:"true"`
	DbName  string `yaml:"dbName" env-required:"true"`
	SslMode string `yaml:"sslmode" env-default:"disable"`
}

type Server struct {
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"10s"`
	Address     string        `yaml:"address" env-default:"localhost:8080"`
}

const op = "config.MustLoad: "

func MustLoad() *Config {
	if err := godotenv.Load("local.env"); err != nil {
		slog.Error(op, err)
		os.Exit(400)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		slog.Error(op, slog.String("env", "is empty"))
		os.Exit(404)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Error(op, slog.String("config", "doesn't exist"))
		os.Exit(404)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		slog.Error(op, slog.String("path", "couldn't find config path"))
		os.Exit(400)
	}
	return &cfg
}
