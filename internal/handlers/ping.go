package handlers

import (
	"net/http"

	"github.com/Wrestler094/shortener/internal/storage"
)

// PingHandler обрабатывает запросы на проверку доступности хранилища
type PingHandler struct {
	storage storage.IStorage // Хранилище для проверки доступности
}

// NewPingHandler создает новый экземпляр PingHandler
// storage - хранилище для проверки доступности
func NewPingHandler(storage storage.IStorage) *PingHandler {
	return &PingHandler{storage: storage}
}

// Ping обрабатывает GET-запрос для проверки доступности хранилища
// Если хранилище поддерживает интерфейс IPingableStorage, проверяет его доступность
// Возможные коды ответа:
// - 200: Хранилище доступно
// - 500: Хранилище недоступно
func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if ps, ok := h.storage.(storage.IPingableStorage); ok {
		if err := ps.Ping(r.Context()); err != nil {
			http.Error(w, "DB unavailable", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
