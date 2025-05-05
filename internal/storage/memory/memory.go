package memory

import (
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
// shortURL - сокращенный URL
// originalURL - оригинальный URL
// userID - идентификатор пользователя (не используется в памяти)
func (ms *MemoryStorage) Save(shortURL string, originalURL string, _ string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.storage[shortURL] = originalURL
	return nil
}

// SaveBatch сохраняет пакет URL в хранилище
// batch - карта сокращенных URL к оригинальным URL
// userID - идентификатор пользователя (не используется в памяти)
func (ms *MemoryStorage) SaveBatch(batch map[string]string, _ string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	for k, v := range batch {
		ms.storage[k] = v
	}

	return nil
}

// Get возвращает оригинальный URL по сокращенному
// shortURL - сокращенный URL
// Возвращает:
// - оригинальный URL
// - флаг удаления (всегда false для памяти)
// - флаг наличия URL в хранилище
func (ms *MemoryStorage) Get(shortURL string) (string, bool, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	url, ok := ms.storage[shortURL]
	return url, false, ok
}

// GetUserURLs возвращает список URL пользователя
// uuid - идентификатор пользователя
// В памяти всегда возвращает пустой список
func (ms *MemoryStorage) GetUserURLs(uuid string) ([]dto.UserURLItem, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	return make([]dto.UserURLItem, 0), nil
}

// DeleteUserURLs помечает URL пользователя как удаленные
// В памяти не реализовано
func (ms *MemoryStorage) DeleteUserURLs(_ string, _ []string) error {
	return nil
}

// FindShortByOriginalURL ищет сокращенный URL по оригинальному
// originalURL - оригинальный URL
// Возвращает сокращенный URL или ошибку, если URL не найден
func (ms *MemoryStorage) FindShortByOriginalURL(originalURL string) (string, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	for short, orig := range ms.storage {
		if orig == originalURL {
			return short, nil
		}
	}

	return "", fmt.Errorf("could not find original url")
}
