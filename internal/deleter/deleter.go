// Package deleter предоставляет реализацию асинхронного удаления сокращённых URL.
// Основная задача — накопление URL для удаления в буфере и их последующая пакетная очистка
// через заданный интервал времени в фоновом процессе.
//
// Это позволяет разгрузить основную логику обработки запросов и минимизировать количество операций с базой данных.
package deleter

import (
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/storage"
)

// Deleter определяет интерфейс для асинхронного удаления URL
type Deleter interface {
	// QueueForDeletion добавляет URL в очередь на удаление
	QueueForDeletion(shortID, userID string)
	// StartBackgroundFlusher запускает фоновый процесс периодического удаления URL
	StartBackgroundFlusher()
}

// URLDeleter реализует асинхронное удаление URL
type URLDeleter struct {
	mu       sync.Mutex          // Мьютекс для синхронизации доступа к буферу
	buffer   map[string][]string // Буфер URL для удаления: map[userID][]shortID
	storage  storage.IStorage    // Хранилище для удаления URL
	interval time.Duration       // Интервал между попытками удаления
}

// NewURLDeleter создает новый экземпляр URLDeleter
// storage - хранилище для удаления URL
// interval - интервал между попытками удаления
func NewURLDeleter(storage storage.IStorage, interval time.Duration) *URLDeleter {
	return &URLDeleter{
		buffer:   make(map[string][]string),
		storage:  storage,
		interval: interval,
	}
}

// QueueForDeletion добавляет URL в очередь на удаление
// shortID - сокращенный URL для удаления
// userID - идентификатор пользователя
func (d *URLDeleter) QueueForDeletion(shortID, userID string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.buffer[userID] = append(d.buffer[userID], shortID)
}

// StartBackgroundFlusher запускает фоновый процесс периодического удаления URL
// Процесс запускается в отдельной горутине и периодически пытается удалить URL из буфера
func (d *URLDeleter) StartBackgroundFlusher() {
	ticker := time.NewTicker(d.interval)

	go func() {
		for range ticker.C {
			d.flush()
		}
	}()
}

// flush пытается удалить все URL из буфера
// При ошибке удаления URL остаются в буфере для следующей попытки
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
