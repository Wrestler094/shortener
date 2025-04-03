package memory

import "sync"

var (
	// map[shortenURL]URL
	storage = make(map[string]string)
	mu      sync.RWMutex
)

func Save(shortenURL string, url string) {
	mu.Lock()
	storage[shortenURL] = url
	mu.Unlock()
}

func Get(shortenURL string) (string, bool) {
	mu.RLock()
	url, ok := storage[shortenURL]
	mu.RUnlock()
	return url, ok
}
