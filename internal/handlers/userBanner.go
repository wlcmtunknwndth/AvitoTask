package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/wlcmtunknwndth/AvitoTask/internal/auth"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpResponse"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogAttr"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
	"io"
	"log/slog"
	"net/http"
)

type queryUserBanner struct {
	TagId           uint `json:"tag_id"`
	FeatureId       uint `json:"feature_id"`
	UseLastRevision bool `json:"use_last_revision,omitempty"`
}

func (h *Handler) UserBanner(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.userBanner.UserBanner"

	info, err := auth.Access(w, r)
	if err != nil {
		slog.Error("couldn't handle access token: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusUnauthorized, statusUnauthorized)
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
		httpResponse.WriteResponse(w, http.StatusInternalServerError, statusInternalServerError)
		return
	}

	var qry queryUserBanner
	err = json.Unmarshal(body, &qry)
	if err != nil {
		slog.Error("couldn't unmarshall body: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusInternalServerError, statusInternalServerError)
		return
	}

	var banner *storage.Banner
	if qry.UseLastRevision || info.IsAdmin {
		banner, err = h.db.GetBanner(qry.FeatureId, qry.TagId)
		if err != nil {
			slog.Error("couldn't find banner: ", slogAttr.OpInfo(op), slogAttr.Err(err))
			httpResponse.WriteResponse(w, http.StatusNotFound, statusNotFound)
			return
		}
		h.cacher.CacheOrder(*banner)

	} else {
		var ok bool
		banner, ok = h.cacher.GetOrder(fmt.Sprintf("%d/%d", qry.FeatureId, qry.TagId))
		if !ok {
			slog.Error("couldn't get cache", slogAttr.OpInfo(op))
			banner, err = h.db.GetBanner(qry.FeatureId, qry.TagId)
			if err != nil {
				slog.Error("couldn't find banner: ", slogAttr.OpInfo(op), slogAttr.Err(err))
				httpResponse.WriteResponse(w, http.StatusNotFound, statusNotFound)
				return
			}
			h.cacher.CacheOrder(*banner)
		} else {
			slog.Info("sent cached order")
		}
	}

	h.WriteBanner(w, banner)
	httpResponse.WriteResponse(w, http.StatusOK, statusOK)
}
