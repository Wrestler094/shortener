package storage

import "sync"

var (
	// map[ShortenUrl]Url
	storage = make(map[string]string)
	mu      sync.RWMutex
)

func Save(shortenUrl string, url string) {
	mu.Lock()
	storage[shortenUrl] = url
	mu.Unlock()
}

func Get(shortenUrl string) (string, bool) {
	mu.RLock()
	url, ok := storage[shortenUrl]
	mu.RUnlock()
	return url, ok
}
