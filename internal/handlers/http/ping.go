package httphandlers

import (
	"net/http"

	"github.com/Wrestler094/shortener/internal/services"
)

// PingHandler обрабатывает запросы на проверку доступности хранилища
type PingHandler struct {
	service *services.PingService
}

// NewPingHandler создает новый экземпляр PingHandler
// storage - хранилище для проверки доступности
func NewPingHandler(service *services.PingService) *PingHandler {
	return &PingHandler{service: service}
}

// Ping обрабатывает GET-запрос для проверки доступности хранилища
// Если хранилище поддерживает интерфейс IPingableStorage, проверяет его доступность
// Возможные коды ответа:
// - 200: Хранилище доступно
// - 500: Хранилище недоступно
func (ph *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	err := ph.service.Ping(r.Context())
	if err != nil {
		http.Error(w, "DB unavailable", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
