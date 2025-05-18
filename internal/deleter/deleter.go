package deleter

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/storage"
)

// Deleter определяет интерфейс для асинхронного удаления URL
type Deleter interface {
	QueueForDeletion(shortID, userID string)
	StartBackgroundFlusher()
	Stop()
}

// URLDeleter реализует асинхронное удаление URL
type URLDeleter struct {
	mu       sync.Mutex
	buffer   map[string][]string
	storage  storage.IStorage
	interval time.Duration

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewURLDeleter создаёт новый экземпляр URLDeleter
func NewURLDeleter(storage storage.IStorage, interval time.Duration) *URLDeleter {
	ctx, cancel := context.WithCancel(context.Background())
	return &URLDeleter{
		buffer:   make(map[string][]string),
		storage:  storage,
		interval: interval,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// QueueForDeletion добавляет URL в очередь на удаление
func (d *URLDeleter) QueueForDeletion(shortID, userID string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.buffer[userID] = append(d.buffer[userID], shortID)
}

// StartBackgroundFlusher запускает фоновый процесс удаления
func (d *URLDeleter) StartBackgroundFlusher() {
	d.wg.Add(1)
	ticker := time.NewTicker(d.interval)

	go func() {
		defer d.wg.Done()
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				d.flush()
			case <-d.ctx.Done():
				logger.Log.Info("Deleter received shutdown signal, flushing...")
				d.flush() // при остановке flush в последний раз
				return
			}
		}
	}()
}

// Stop завершает фоновые процессы и ждёт завершения
func (d *URLDeleter) Stop() {
	d.cancel()
	d.wg.Wait()
}

// flush удаляет URL из буфера
func (d *URLDeleter) flush() {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Создаем контекст с таймаутом в 5 секунд для операции удаления
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for userID, shortIDs := range d.buffer {
		if len(shortIDs) == 0 {
			continue
		}
		err := d.storage.DeleteUserURLs(ctx, userID, shortIDs)
		if err != nil {
			logger.Log.Error("Batch delete failed",
				zap.String("user_id", userID),
				zap.Error(err),
				zap.Error(ctx.Err()))
			continue
		}
		delete(d.buffer, userID)
	}
}
