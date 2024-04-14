package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogAttr"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	DbConfig DbConfig `yaml:"dbConfig" env-required:"true"`
	Server   Server   `yaml:"server"`
}

type DbConfig struct {
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
	const op = "config.MustLoad"
	//if err := godotenv.Load("local.env"); err != nil {
	//	slog.Error(op, err)
	//	os.Exit(400)
	//}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		slog.Error("couldn't load config path", slogAttr.OpInfo(op))
		os.Exit(404)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Error("couldn't find config file", slogAttr.OpInfo(op), slogAttr.Err(err))
		os.Exit(404)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		slog.Error("couldn't read config file", slogAttr.OpInfo(op), slogAttr.Err(err))
		os.Exit(400)
	}
	return &cfg
}
