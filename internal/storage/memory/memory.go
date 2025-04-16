package memory

import (
	"sync"
)

// TODO: Попробовать переписать с type shortURL / OriginalURL

type MemoryStorage struct {
	// map[shortURL]originalURL
	storage map[string]string
	mu      sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	// TODO: Realise urls recover from file

	return &MemoryStorage{
		storage: make(map[string]string),
		mu:      sync.RWMutex{},
	}
}

func (ms *MemoryStorage) Save(shortURL string, OriginalURL string) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.storage[shortURL] = OriginalURL
}

func (ms *MemoryStorage) Get(shortURL string) (string, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	url, ok := ms.storage[shortURL]
	return url, ok
}
