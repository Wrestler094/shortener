// Реализация хранения URL в файловой системе.
// Обеспечивает персистентное хранение пар сокращенных и оригинальных URL
// в JSON-формате с поддержкой:
// - Атомарной записи новых URL
// - Восстановления данных при перезапуске
// - Потокобезопасного доступа к файлу
// - Логирования ошибок операций с файлом
package persistence

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/logger"
)

// FileStorage реализует хранение URL в файле
type FileStorage struct {
	path string     // Путь к файлу для хранения URL
	mu   sync.Mutex // Мьютекс для синхронизации доступа к файлу
}

// NewFileStorage создает новый экземпляр FileStorage
// path - путь к файлу для хранения URL
func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
}

// SaveURL сохраняет пару сокращенный URL - оригинальный URL в файл
// shortURL - сокращенный URL
// originalURL - оригинальный URL
// Если путь к файлу не указан, операция игнорируется
func (fs *FileStorage) SaveURL(shortURL, originalURL string) {
	if fs.path == "" {
		return
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	data := fmt.Sprintf("{\"short_url\":\"%s\",\"original_url\":\"%s\"}\n", shortURL, originalURL)
	byteSlice := []byte(data)

	file, err := os.OpenFile(fs.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Log.Error("Error of open file", zap.Error(err))
		return
	}

	_, err = file.Write(byteSlice)
	if err != nil {
		logger.Log.Error("Error of writing url to file", zap.Error(err))
		return
	}

	err = file.Close()
	if err != nil {
		logger.Log.Error("Error of closing file", zap.Error(err))
		return
	}
}

// URLEntry представляет запись URL в файле
type URLEntry struct {
	ShortURL    string `json:"short_url"`    // Сокращенный URL
	OriginalURL string `json:"original_url"` // Оригинальный URL
}

// RecoverURLs восстанавливает URL из файла
// Возвращает карту сокращенных URL к оригинальным URL
// Если путь к файлу не указан или произошла ошибка, возвращает пустую карту
func (fs *FileStorage) RecoverURLs() map[string]string {
	result := make(map[string]string)

	if fs.path == "" {
		return result
	}

	file, err := os.OpenFile(fs.path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Log.Error("Error opening file", zap.Error(err))
		return result
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Log.Error("Error closing file", zap.Error(err))
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var entry URLEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			logger.Log.Error("Error unmarshalling line", zap.Error(err))
			continue
		}

		result[entry.ShortURL] = entry.OriginalURL
	}

	if err := scanner.Err(); err != nil {
		logger.Log.Error("Scanner error", zap.Error(err))
	}

	return result
}
