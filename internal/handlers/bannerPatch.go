package handlers

import (
	"encoding/json"
	"github.com/wlcmtunknwndth/AvitoTask/internal/auth"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpResponse"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogAttr"
	"github.com/wlcmtunknwndth/AvitoTask/internal/storage"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) BannerPatch(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.BannerPatch"

	info, err := auth.Access(w, r)
	if err != nil {
		slog.Error("couldn't handle access token: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusUnauthorized, statusUnauthorized)
		return
	}
	if !info.IsAdmin {
		httpResponse.WriteResponse(w, http.StatusForbidden, statusNoAccess)
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
		httpResponse.WriteResponse(w, http.StatusBadRequest, statusBadRequest)
		return
	}

	var banner storage.Banner

	err = json.Unmarshal(body, &banner)
	if err != nil {
		slog.Error("couldn't unmarshal body: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusBadRequest, statusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		slog.Error("couldn't get id from url path: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusBadRequest, statusBadRequest)
		return
	}

	err = h.db.UpdateBannerById(uint(id), &banner)
	if err != nil {
		slog.Error("couldn't update banner: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusNotFound, statusNotFound)
		return
	}

	httpResponse.WriteResponse(w, http.StatusOK, statusOK)
}
