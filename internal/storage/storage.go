package storage

import (
	"context"

	"github.com/Wrestler094/shortener/internal/dto"
)

type IStorage interface {
	Save(string, string) error
	Get(string) (string, bool)
	GetUserURLs(string) ([]dto.UserURLItem, error)
	SaveBatch(map[string]string) error
	FindShortByOriginalURL(string) (string, error)
}

type IPingableStorage interface {
	Ping(ctx context.Context) error
}
