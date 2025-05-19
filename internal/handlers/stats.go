package handlers

import (
	"net"
	"net/http"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/utils"
)

// StatsHandler обрабатывает запросы, связанные со статистикой сервиса
type StatsHandler struct {
	service *services.StatsService
}

// NewStatsHandler создает новый экземпляр StatsHandler
func NewStatsHandler(service *services.StatsService) *StatsHandler {
	return &StatsHandler{
		service: service,
	}
}

// statsResponse представляет структуру ответа для статистики
type statsResponse struct {
	URLs  int `json:"urls"`  // Количество URL в системе
	Users int `json:"users"` // Количество пользователей в системе
}

// GetStats обрабатывает GET запрос для получения статистики сервиса
// Возвращает количество URL и пользователей в системе
// Доступ разрешен только с доверенных IP-адресов
func (sh *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	if !sh.isTrusted(r) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	urls, users, err := sh.service.GetStats(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, statsResponse{URLs: urls, Users: users})
}

// isTrusted проверяет, является ли IP-адрес запроса доверенным
// Проверка осуществляется на основе настройки FlagTrustedSubnet
func (sh *StatsHandler) isTrusted(r *http.Request) bool {
	if configs.FlagTrustedSubnet == "" {
		return false
	}

	ipStr := r.Header.Get("X-Real-IP")
	ip := net.ParseIP(ipStr)
	_, subnet, err := net.ParseCIDR(configs.FlagTrustedSubnet)
	if err != nil || ip == nil {
		return false
	}

	return subnet.Contains(ip)
}
