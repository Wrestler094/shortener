package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/Wrestler094/shortener/internal/dto"
)

// MemoryStorage реализует хранилище URL в памяти с поддержкой конкурентного доступа
type MemoryStorage struct {
	storage map[string]string // map[shortURL]originalURL - хранилище URL
	mu      sync.RWMutex      // Мьютекс для синхронизации доступа
}

// NewMemoryStorage создает новый экземпляр хранилища в памяти
// recoveredUrls - карта восстановленных URL при запуске приложения
func NewMemoryStorage(recoveredUrls map[string]string) *MemoryStorage {
	return &MemoryStorage{storage: recoveredUrls}
}

// Save сохраняет пару сокращенный URL - оригинальный URL в хранилище
// ctx - контекст запроса
// shortURL - сокращенный URL
// originalURL - оригинальный URL
// userID - идентификатор пользователя (не используется в памяти)
func (ms *MemoryStorage) Save(_ context.Context, shortURL string, originalURL string, _ string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.storage[shortURL] = originalURL
	return nil
}

// SaveBatch сохраняет пакет URL в хранилище
// ctx - контекст запроса
// batch - карта сокращенных URL к оригинальным URL
// userID - идентификатор пользователя (не используется в памяти)
func (ms *MemoryStorage) SaveBatch(_ context.Context, batch map[string]string, _ string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	for k, v := range batch {
		ms.storage[k] = v
	}

	return nil
}

// Get возвращает оригинальный URL по сокращенному
// ctx - контекст запроса
// shortURL - сокращенный URL
// Возвращает:
// - оригинальный URL
// - флаг удаления (всегда false для памяти)
// - флаг наличия URL в хранилище
func (ms *MemoryStorage) Get(_ context.Context, shortURL string) (string, bool, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	url, ok := ms.storage[shortURL]
	return url, false, ok
}

// GetUserURLs возвращает список URL пользователя
// ctx - контекст запроса
// uuid - идентификатор пользователя
// В памяти всегда возвращает пустой список
func (ms *MemoryStorage) GetUserURLs(_ context.Context, uuid string) ([]dto.UserURLItem, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	return make([]dto.UserURLItem, 0), nil
}

// DeleteUserURLs помечает URL пользователя как удаленные
// ctx - контекст запроса
// В памяти не реализовано
func (ms *MemoryStorage) DeleteUserURLs(_ context.Context, _ string, _ []string) error {
	return nil
}

// FindShortByOriginalURL ищет сокращенный URL по оригинальному
// ctx - контекст запроса
// originalURL - оригинальный URL
// Возвращает сокращенный URL или ошибку, если URL не найден
func (ms *MemoryStorage) FindShortByOriginalURL(_ context.Context, originalURL string) (string, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	for short, orig := range ms.storage {
		if orig == originalURL {
			return short, nil
		}
	}

	return "", fmt.Errorf("could not find original url")
}
