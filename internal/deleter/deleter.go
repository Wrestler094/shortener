package deleter

import (
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/storage"
)

type Deleter interface {
	QueueForDeletion(shortID, userID string)
	StartBackgroundFlusher()
}

type URLDeleter struct {
	mu       sync.Mutex
	buffer   map[string][]string // map[userID][]shortID
	storage  storage.IStorage
	interval time.Duration
}

func NewURLDeleter(storage storage.IStorage, interval time.Duration) *URLDeleter {
	return &URLDeleter{
		buffer:   make(map[string][]string),
		storage:  storage,
		interval: interval,
	}
}

func (d *URLDeleter) QueueForDeletion(shortID, userID string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.buffer[userID] = append(d.buffer[userID], shortID)
}

func (d *URLDeleter) StartBackgroundFlusher() {
	ticker := time.NewTicker(d.interval)

	go func() {
		for range ticker.C {
			d.flush()
		}
	}()
}

func (d *URLDeleter) flush() {
	d.mu.Lock()
	defer d.mu.Unlock()

	for userID, shortIDs := range d.buffer {
		if len(shortIDs) == 0 {
			continue
		}

		err := d.storage.DeleteUserURLs(userID, shortIDs)
		if err != nil {
			logger.Log.Error("Batch delete failed", zap.String("user_id", userID), zap.Error(err))
			continue
		}
		delete(d.buffer, userID)
	}
}
