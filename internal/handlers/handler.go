package handlers

import (
	"encoding/json"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpErrors"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpResponse"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
	"net/http"
)

const (
	statusOK                  = "OK"
	statusDeleted             = "Banner has been successfully deleted" // 204
	statusBadRequest          = "Bad request"                          // 400
	statusNotFound            = "Banner wasn't found"                  // 400
	statusUnauthorized        = "User unauthorized"                    // 401
	statusNoAccess            = "User has no access"                   // 403
	statusInternalServerError = "Internal server error"                // 500
)

type Storage interface {
	GetPassword(login string) (string, error)
	//RegisterUser(usr *auth.User) error
	//DeleteUser(login string) error

	SaveBanner(banner *storage.Banner) error
	DeleteBanner(id uint) error
	GetBanner(featureId, tag uint) (*storage.Banner, error)
	GetBannersByFeature(featureId uint) ([]storage.Banner, error)
	GetBannersByTag(tag uint) ([]storage.Banner, error)
	UpdateBannerById(id uint, banner *storage.Banner) error
}

type Cacher interface {
	CacheOrder(banner storage.Banner)
	GetOrder(uuid string) (*storage.Banner, bool)
	Restore() error
	SaveCache() error
}

type Handler struct {
	db     Storage
	cacher Cacher
}

func NewHandler(db Storage, cacher Cacher) *Handler {
	return &Handler{
		db:     db,
		cacher: cacher,
	}
}

func (h *Handler) WriteBanner(w http.ResponseWriter, banner *storage.Banner) {
	const op = "handler.WriteBanner"
	data, err := json.Marshal(banner)
	if err != nil {
		httpResponse.WriteResponse(w, http.StatusNotFound, httpErrors.Error404)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}
}

func (h *Handler) WriteBanners(w http.ResponseWriter, banners []storage.Banner) {
	const op = "handler.WriteBanners"
	data, err := json.Marshal(banners)
	if err != nil {
		httpResponse.WriteResponse(w, http.StatusNotFound, httpErrors.Error404)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}
}
