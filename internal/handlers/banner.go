package handlers

import (
	"encoding/json"
	"github.com/wlcmtunknwndth/AvitoTask/internal/auth"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpErrors"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpResponse"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogAttr"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
	"io"
	"log/slog"
	"net/http"
)

type BannerGet struct {
	TagId     uint `json:"tag_id,omitempty"`
	FeatureId uint `json:"feature_id,omitempty"`
	Limit     uint `json:"limit,omitempty"`
	Offset    uint `json:"offset,omitempty"`
}

func (h *Handler) BannerGet(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.banner.BannerGet"

	info, err := auth.Access(w, r)
	if err != nil || !info.IsAdmin {
		slog.Error("couldn't handle access token: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusUnauthorized, httpErrors.Error401)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("couldn't close request body: ", slogAttr.OpInfo(op), slogAttr.Err(err))
			return
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error decoding request: ", err)
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}

	var qry BannerGet
	err = json.Unmarshal(body, &qry)
	if err != nil {
		slog.Error("couldn't unmarshall body: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}

	if qry.TagId != 0 && qry.FeatureId != 0 {
		banner, err := h.db.GetBanner(qry.FeatureId, qry.TagId)
		if err != nil {
			slog.Error("couldn't get banner: ", slogAttr.OpInfo(op), slogAttr.Err(err))
			httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
			return
		}
		h.cacher.CacheOrder(*banner)
		h.WriteBanner(w, banner)
	} else if qry.FeatureId != 0 {
		banners, err := h.db.GetBannersByFeature(qry.FeatureId)
		if err != nil {
			slog.Error("couldn't get banners by feature: ", slogAttr.OpInfo(op), slogAttr.Err(err))
			httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
			return
		}
		h.WriteBanners(w, banners)

	} else if qry.TagId != 0 {
		banners, err := h.db.GetBannersByTag(qry.TagId)
		if err != nil {
			slog.Error("couldn't get banners by feature: ", slogAttr.OpInfo(op), slogAttr.Err(err))
			httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
			return
		}
		h.WriteBanners(w, banners)
	} else {
		httpResponse.WriteResponse(w, http.StatusBadRequest, httpErrors.Error400)
	}

}

func (h *Handler) BannerPost(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.banner.BannerPost"

	info, err := auth.Access(w, r)
	if err != nil || !info.IsAdmin {
		slog.Error("couldn't handle access token: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusUnauthorized, httpErrors.Error401)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("couldn't close request body: ", slogAttr.OpInfo(op), slogAttr.Err(err))
			return
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error decoding request: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}

	var banner storage.Banner
	err = json.Unmarshal(body, &banner)
	if err != nil {
		slog.Error("error unmarshalling request:", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}

	err = h.db.SaveBanner(&banner)
	if err != nil {
		slog.Error("error unmarshalling request:", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusInternalServerError, httpErrors.Error500)
		return
	}
	h.cacher.CacheOrder(banner)

}
