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
	return &MemoryStorage{
		storage: make(map[string]string),
		mu:      sync.RWMutex{},
	}
}

func (m *MemoryStorage) Save(shortURL string, OriginalURL string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.storage[shortURL] = OriginalURL
}

func (m *MemoryStorage) Get(shortURL string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	url, ok := m.storage[shortURL]
	return url, ok
}
