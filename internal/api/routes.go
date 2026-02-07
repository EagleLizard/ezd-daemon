package api

import (
	"encoding/json"
	"net/http"

	"github.com/EagleLizard/ezd-daemon/internal/api/ctrl"
	"github.com/EagleLizard/ezd-daemon/internal/lib/config"
)

func addRoutes(mux *http.ServeMux, cfg *config.EzdDConfigType) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})
	mux.HandleFunc("POST /v1/hook/gh", ctrl.PostGhHook(cfg))
}
