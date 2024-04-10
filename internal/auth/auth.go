package auth

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type Storage interface {
	GetPass(ctx context.Context, username string) (string, error)
	RegisterUser(ctx context.Context, user User) error
}

type Auth struct {
	Db Storage
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var usr User
	var ctx context.Context

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		slog.Info("couldn't process request: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = a.Db.RegisterUser(ctx, usr); err != nil {
		slog.Error("couldn't register new user: ", err)
	}

	WriteNewToken(w, usr, AccessToken)
	WriteNewToken(w, usr, RefreshToken)
}

func (a *Auth) LogIn(w http.ResponseWriter, r *http.Request) {
	var usr User
	var ctx context.Context

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass, err := a.Db.GetPass(ctx, usr.Username)
	if err != nil || pass != usr.Password {
		slog.Error("pass is not valid or couldn't check password:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	WriteNewToken(w, usr, AccessToken)
	WriteNewToken(w, usr, RefreshToken)
}

func (a *Auth) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    AccessToken,
		Expires: time.Now(),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    RefreshToken,
		Expires: time.Now(),
	})
}
