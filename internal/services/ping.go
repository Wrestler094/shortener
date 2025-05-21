package services

import (
	"context"

	"github.com/Wrestler094/shortener/internal/storage"
)

// PingService представляет сервис для проверки доступности хранилища
type PingService struct {
	storage storage.IStorage
}

// NewPingService создает новый экземпляр PingService
func NewPingService(storage storage.IStorage) *PingService {
	return &PingService{
		storage: storage,
	}
}

// Ping проверяет доступность хранилища
func (s *PingService) Ping(ctx context.Context) error {
	if ps, ok := s.storage.(storage.IPingableStorage); ok {
		if err := ps.Ping(ctx); err != nil {
			return err
		}
	}

	return nil
}
