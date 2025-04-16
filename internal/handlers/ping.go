package handlers

import (
	"net/http"

	"github.com/Wrestler094/shortener/internal/storage/postgres"
)

// todo
func HandlePing(w http.ResponseWriter, r *http.Request) {
	if err := postgres.DB.PingContext(r.Context()); err != nil {
		http.Error(w, "database is not available", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
