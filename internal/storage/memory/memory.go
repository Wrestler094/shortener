package memory

import (
	"fmt"
	"sync"

	"github.com/Wrestler094/shortener/internal/dto"
)

type MemoryStorage struct {
	storage map[string]string // map[shortURL]originalURL
	mu      sync.RWMutex
}

func NewMemoryStorage(recoveredUrls map[string]string) *MemoryStorage {
	return &MemoryStorage{storage: recoveredUrls}
}

func (ms *MemoryStorage) Save(shortURL string, originalURL string, _ string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.storage[shortURL] = originalURL
	return nil
}

func (ms *MemoryStorage) SaveBatch(batch map[string]string, _ string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	for k, v := range batch {
		ms.storage[k] = v
	}

	return nil
}

func (ms *MemoryStorage) Get(shortURL string) (string, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	url, ok := ms.storage[shortURL]
	return url, ok
}

func (ms *MemoryStorage) GetUserURLs(uuid string) ([]dto.UserURLItem, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	fmt.Println(uuid)
	return make([]dto.UserURLItem, 0), nil
}

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
