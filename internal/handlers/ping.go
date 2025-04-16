package handlers

import (
	"net/http"

	"github.com/Wrestler094/shortener/internal/storage"
)

type PingHandler struct {
	storage storage.IStorage
}

func NewPingHandler(storage storage.IStorage) *PingHandler {
	return &PingHandler{storage: storage}
}

func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if ps, ok := h.storage.(storage.IPingableStorage); ok {
		if err := ps.Ping(r.Context()); err != nil {
			http.Error(w, "DB unavailable", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
