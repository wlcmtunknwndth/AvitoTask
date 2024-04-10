package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wlcmtunknwndth/AvitoTask/internal/config"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogError"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage/postgres"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	pgsql, err := postgres.New(cfg.DbConfig)
	if err != nil {
		slog.Error("couldn't run postgres ", slogError.Err(err))
	}
	defer func(storage *postgres.Storage) {
		err = storage.Close()
		if err != nil {
			slog.Error("couldn't close postgres: ", slogError.Err(err))
		}
	}(pgsql)

	router := chi.NewRouter()
	router.Use(middleware.RequestID) // adds requestID to logs
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat) // adds request format
	router.Use(middleware.Logger)

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	if err = srv.ListenAndServe(); err != nil {
		slog.Error("failed to run server: ", slogError.Err(err))
	}
	slog.Info("application finished")
}
