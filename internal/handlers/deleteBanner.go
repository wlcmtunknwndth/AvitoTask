package handlers

import (
	"encoding/json"
	"github.com/wlcmtunknwndth/AvitoTask/internal/auth"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/httpResponse"
	"github.com/wlcmtunknwndth/AvitoTask/internal/lib/slogAttr"
	"io"
	"log/slog"
	"net/http"
)

type qryDelete struct {
	Id uint `json:"id"`
}

func (h *Handler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.DeleteBanner"

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

	var qry qryDelete

	if err = json.Unmarshal(body, &qry); err != nil {
		slog.Error("error unmarshalling request: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusBadRequest, statusBadRequest)
		return
	}

	if err = h.db.DeleteBanner(qry.Id); err != nil {
		slog.Error("error unmarshalling request: ", slogAttr.OpInfo(op), slogAttr.Err(err))
		httpResponse.WriteResponse(w, http.StatusInternalServerError, statusInternalServerError)
		return
	}

	httpResponse.WriteResponse(w, http.StatusNoContent, statusDeleted)
}
