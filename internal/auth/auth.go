package auth

import (
	"encoding/json"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogAttr"
	"log/slog"
	"net/http"
	"time"
)

type Storage interface {
	GetPassword(string) (string, error)
	RegisterUser(*User) error
	IsAdmin(string) bool
}

type User struct {
	admin    bool   //`json:"isAdmin"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func (u *User) IsAdmin() bool {
	if u.admin {
		return true
	}
	return false
}

type Auth struct {
	Db Storage
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var usr User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		slog.Error("couldn't process request: ", slogAttr.Err(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = a.Db.RegisterUser(&usr); err != nil {
		slog.Error("couldn't register new user: ", slogAttr.Err(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	WriteNewToken(w, usr, AccessToken)
}

func (a *Auth) LogIn(w http.ResponseWriter, r *http.Request) {
	var usr User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass, err := a.Db.GetPassword(usr.Username)
	if err != nil || pass != usr.Password {
		//slog.Error("pass is not valid or couldn't check password:", slogAttr.Err(err))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usr.admin = a.Db.IsAdmin(usr.Username)

	WriteNewToken(w, usr, AccessToken)
	//WriteNewToken(w, usr, RefreshToken)
}

func (a *Auth) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    AccessToken,
		Expires: time.Now(),
	})
}
