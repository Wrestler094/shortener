package services

import (
	"context"

	"github.com/Wrestler094/shortener/internal/storage"
)

// StatsService предоставляет методы для работы со статистикой сервиса сокращения URL
type StatsService struct {
	storage storage.IStatsStorage
}

// NewStatsService создает новый экземпляр StatsService
func NewStatsService(storage storage.IStatsStorage) *StatsService {
	return &StatsService{
		storage: storage,
	}
}

// GetStats возвращает общую статистику сервиса:
// - количество сокращенных URL
// - количество пользователей
func (s *StatsService) GetStats(ctx context.Context) (urls int, users int, err error) {
	return s.storage.GetStats(ctx)
}
