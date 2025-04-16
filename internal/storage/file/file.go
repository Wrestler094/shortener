package file

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/logger"
)

var (
	mu sync.Mutex
)

type URLEntry struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func SaveURL(shortURL string, originalURL string) {
	mu.Lock()
	defer mu.Unlock()

	data := fmt.Sprintf("{\"short_url\":\"%s\",\"original_url\":\"%s\"}\n", shortURL, originalURL)
	byteSlice := []byte(data)

	file, err := os.OpenFile(configs.FlagFileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
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

//
//func RecoverURLs() {
//	file, err := os.OpenFile(configs.FlagFileStoragePath, os.O_RDONLY|os.O_CREATE, 0666)
//	if err != nil {
//		logger.Log.Error("Error of writing url to file", zap.Error(err))
//		return
//	}
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		line := scanner.Text()
//
//		var entry URLEntry
//		if err := json.Unmarshal([]byte(line), &entry); err != nil {
//			logger.Log.Error("Error of reading url from file", zap.Error(err))
//			continue
//		}
//
//		storage.Storage.Save(entry.ShortURL, entry.OriginalURL)
//	}
//
//	if err := scanner.Err(); err != nil {
//		logger.Log.Error("Scanner error", zap.Error(err))
//		return
//	}
//
//	err = file.Close()
//	if err != nil {
//		logger.Log.Error("Error of closing file", zap.Error(err))
//		return
//	}
//}
