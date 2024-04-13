package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpErrors"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpResponse"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var Key, _ = os.LookupEnv("secret_key")

type Info struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

// TimeToLive == Ttl
const (
	AccessToken  = "access"
	RefreshToken = "refresh"

	TtlAccess  = 4
	TtlRefresh = 5
)

func checkRequest(w http.ResponseWriter, r *http.Request, cookieName string) (*Info, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		slog.Error("error looking for cookies: ", err)
		if errors.Is(err, http.ErrNoCookie) {
			w.WriteHeader(http.StatusUnauthorized)
			return nil, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return nil, fmt.Errorf("no cookies found: %s", err)
	}

	tokenStr := cookie.Value

	var info Info

	token, err := jwt.ParseWithClaims(tokenStr, &info, func(token *jwt.Token) (any, error) {
		return []byte(Key), nil
	})

	if err != nil {
		slog.Error("couldn't parse jwt: ", err)
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			w.WriteHeader(http.StatusUnauthorized)
			return nil, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return nil, fmt.Errorf("auth error: %s", err)
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, fmt.Errorf("the token is invalid")
	}
	return &info, nil
}

func Access(w http.ResponseWriter, r *http.Request) (*Info, error) {
	const op = "auth.jwtAuth.Access"
	info, err := checkRequest(w, r, AccessToken)
	if err != nil {
		httpResponse.WriteResponse(w, http.StatusUnauthorized, httpErrors.Error401)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return info, nil
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	info, err := checkRequest(w, r, AccessToken)
	if err != nil {
		return
	}

	// gives a new access token only if previous is about to die in 30 secs
	//if time.Until(info.ExpiresAt.Time) > 30*time.Second {
	//	slog.Error("token ", RefreshToken, " is about to out of clock")
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	expiresAt := time.Now().Add(TtlAccess * time.Minute)

	info.ExpiresAt = jwt.NewNumericDate(expiresAt)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, info)

	tokenStr, err := token.SignedString([]byte(Key))
	if err != nil {
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    AccessToken,
		Value:   tokenStr,
		Expires: expiresAt,
	})

}

func WriteNewToken(w http.ResponseWriter, usr User, tokenName string) {
	var expireAt time.Time
	switch tokenName {
	//case RefreshToken:
	//	expireAt = time.Now().Add(TtlRefresh * time.Minute)
	case AccessToken:
		expireAt = time.Now().Add(TtlAccess * time.Minute)
	default:
		return
	}

	inf := &Info{
		Username: usr.Username,
		IsAdmin:  usr.admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, inf)

	tokenStr, err := token.SignedString([]byte(Key))
	if err != nil {
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    tokenName,
		Value:   tokenStr,
		Expires: expireAt,
	})
}

//
//func WriteBothTokens(w http.ResponseWriter, usr *User) {
//	var expireAt time.Time
//	inf := &Info{
//		Username: usr.Username,
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(expireAt),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS512, inf)
//
//	tokenStr, err := token.SignedString([]byte(Key))
//	if err != nil {
//		httpResponse.WriteResponse(w, http.StatusInternalServerError, "couldn't send access tokens")
//		return
//	}
//
//	http.SetCookie(w, &http.Cookie{
//		Name:    tokenName,
//		Value:   tokenStr,
//		Expires: expireAt,
//	})
//
//	//http.Set
//}
