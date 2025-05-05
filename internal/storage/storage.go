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

// IStorage определяет интерфейс для хранения URL
type IStorage interface {
	// Save сохраняет пару сокращенный URL - оригинальный URL
	// shortURL - сокращенный URL
	// originalURL - оригинальный URL
	// userID - идентификатор пользователя
	Save(string, string, string) error

	// SaveBatch сохраняет пакет URL
	// urls - карта сокращенных URL к оригинальным URL
	// userID - идентификатор пользователя
	SaveBatch(map[string]string, string) error

	// Get возвращает оригинальный URL по сокращенному
	// shortURL - сокращенный URL
	// Возвращает:
	// - оригинальный URL
	// - флаг удаления
	// - флаг наличия URL в хранилище
	Get(string) (string, bool, bool)

	// GetUserURLs возвращает список URL пользователя
	// userID - идентификатор пользователя
	GetUserURLs(string) ([]dto.UserURLItem, error)

	// DeleteUserURLs помечает URL пользователя как удаленные
	// userID - идентификатор пользователя
	// urls - список сокращенных URL для удаления
	DeleteUserURLs(string, []string) error

	// FindShortByOriginalURL ищет сокращенный URL по оригинальному
	// originalURL - оригинальный URL
	FindShortByOriginalURL(string) (string, error)
}

// IPingableStorage определяет интерфейс для проверки доступности хранилища
type IPingableStorage interface {
	// Ping проверяет соединение с хранилищем
	// ctx - контекст для выполнения запроса
	Ping(ctx context.Context) error
}
