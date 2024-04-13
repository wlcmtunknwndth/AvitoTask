package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wlcmtunknwndth/AvitoTask/internal/auth"
	"github.com/wlcmtunknwndth/AvitoTask/internal/cacher"
	"github.com/wlcmtunknwndth/AvitoTask/internal/config"
	"github.com/wlcmtunknwndth/AvitoTask/internal/handlers"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogAttr"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage/postgres"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	cfg := config.MustLoad()

	pgsql, err := postgres.New(cfg.DbConfig)
	if err != nil {
		slog.Error("couldn't run postgres ", slogAttr.Err(err))
	}
	defer func(storage *postgres.Storage) {
		err = storage.Close()
		if err != nil {
			slog.Error("couldn't close postgres: ", slogAttr.Err(err))
		}
	}(pgsql)

	authService := auth.Auth{Db: pgsql}

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

	router.Post("/register", authService.Register)
	router.Post("/login", authService.LogIn)
	router.Post("/logout", authService.LogOut)

	cacherHandler := cacher.New(pgsql, 5*time.Minute, 20*time.Minute)
	//err = cacherHandler.Restore()
	err = cacherHandler.Restore()
	if err != nil {
		slog.Error("couldn't restore cache: ", slogAttr.Err(err))
		return
	}
	slog.Info("cache restored")

	//time.NewTimer()
	go func(cache *cacher.Cacher) {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := cache.SaveCache()
				if err != nil {
					continue
				}
			}
		}
	}(cacherHandler)

	httpHandler := handlers.NewHandler(pgsql, cacherHandler)

	router.Get("/user_banner", httpHandler.UserBanner)
	router.Get("/banner", httpHandler.BannerGet)
	router.Post("/banner", httpHandler.BannerPost)
	router.Patch("/banner/{id}", httpHandler.BannerPatch)
	router.Delete("/delete", httpHandler.DeleteBanner)

	if err = srv.ListenAndServe(); err != nil {
		slog.Error("failed to run server: ", slogAttr.Err(err))
	}
	slog.Info("application finished")
}
