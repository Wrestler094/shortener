package handlers

import (
	"net"
	"net/http"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/utils"
)

type StatsHandler struct {
	service *services.StatsService
}

func NewStatsHandler(service *services.StatsService) *StatsHandler {
	return &StatsHandler{
		service: service,
	}
}

type statsResponse struct {
	URLs  int `json:"urls"`
	Users int `json:"users"`
}

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
