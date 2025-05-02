package storage

import (
	"context"

	"github.com/Wrestler094/shortener/internal/dto"
	"github.com/Wrestler094/shortener/internal/storage/memory"
	"github.com/Wrestler094/shortener/internal/storage/postgres"
)

var (
	_ IStorage = (*postgres.PostgresStorage)(nil)
	_ IStorage = (*memory.MemoryStorage)(nil)
)

type IStorage interface {
	Save(string, string, string) error
	SaveBatch(map[string]string, string) error
	Get(string) (string, bool, bool)
	GetUserURLs(string) ([]dto.UserURLItem, error)
	DeleteUserURLs(string, []string) error
	FindShortByOriginalURL(string) (string, error)
}

type IPingableStorage interface {
	Ping(ctx context.Context) error
}
