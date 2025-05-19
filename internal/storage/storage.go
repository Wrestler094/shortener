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

// IStorage объединяет интерфейсы для работы с URL и статистикой хранилища
type IStorage interface {
	// IURLStorage определяет методы для работы с URL в хранилище
	IURLStorage
	// Встраивает интерфейс для получения статистики хранилища
	IStatsStorage
}

// IURLStorage определяет методы для работы с URL в хранилище
type IURLStorage interface {
	// Save сохраняет пару сокращенный URL - оригинальный URL
	// ctx - контекст запроса
	// shortURL - сокращенный URL
	// originalURL - оригинальный URL
	// userID - идентификатор пользователя
	Save(context.Context, string, string, string) error

	// SaveBatch сохраняет пакет URL
	// ctx - контекст запроса
	// urls - карта сокращенных URL к оригинальным URL
	// userID - идентификатор пользователя
	SaveBatch(context.Context, map[string]string, string) error

	// Get возвращает оригинальный URL по сокращенному
	// ctx - контекст запроса
	// shortURL - сокращенный URL
	// Возвращает:
	// - оригинальный URL
	// - флаг удаления
	// - флаг наличия URL в хранилище
	Get(context.Context, string) (string, bool, bool)

	// GetUserURLs возвращает список URL пользователя
	// ctx - контекст запроса
	// userID - идентификатор пользователя
	GetUserURLs(context.Context, string) ([]dto.UserURLItem, error)

	// DeleteUserURLs помечает URL пользователя как удаленные
	// ctx - контекст запроса
	// userID - идентификатор пользователя
	// urls - список сокращенных URL для удаления
	DeleteUserURLs(context.Context, string, []string) error

	// FindShortByOriginalURL ищет сокращенный URL по оригинальному
	// ctx - контекст запроса
	// originalURL - оригинальный URL
	FindShortByOriginalURL(context.Context, string) (string, error)
}

// IStatsStorage определяет методы для получения статистики хранилища
type IStatsStorage interface {
	// GetStats возвращает статистику хранилища
	// ctx - контекст запроса
	// Возвращает:
	// - количество URL в хранилище
	// - количество пользователей
	// - ошибку, если возникла
	GetStats(context.Context) (int, int, error)
}

// IPingableStorage определяет интерфейс для проверки доступности хранилища
type IPingableStorage interface {
	// Ping проверяет соединение с хранилищем
	// ctx - контекст для выполнения запроса
	Ping(context.Context) error
}

// IClosableStorage определяет интерфейс для закрытия соединения с хранилищем
type IClosableStorage interface {
	// Close закрывает соединение с хранилищем
	// Возвращает ошибку в случае неудачного закрытия
	Close() error
}
