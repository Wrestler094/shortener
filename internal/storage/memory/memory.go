package memory

import (
	"sync"
)

type MemoryStorage struct {
	storage map[string]string // map[shortURL]originalURL
	mu      sync.RWMutex
}

func NewMemoryStorage(recoveredUrls map[string]string) *MemoryStorage {
	return &MemoryStorage{storage: recoveredUrls}
}

func (ms *MemoryStorage) Save(shortURL string, originalURL string) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.storage[shortURL] = originalURL
}

func (ms *MemoryStorage) Get(shortURL string) (string, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	url, ok := ms.storage[shortURL]
	return url, ok
}
