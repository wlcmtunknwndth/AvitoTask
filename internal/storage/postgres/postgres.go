package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wlcmtunknwndth/AvitoTask/internal/config"
	"log/slog"
)

type Storage struct {
	db *sql.DB
}

func New(config config.DbConfig) (*Storage, error) {
	const op = "storage.postgres.New"
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.DbUser, config.DbPass, config.DbName, config.SslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		slog.Error(op, slog.String("ping", "error"))
		return nil, err
	} else {
		slog.Info("Pinged db successfully")
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Ping() error {
	return s.db.Ping()
}
