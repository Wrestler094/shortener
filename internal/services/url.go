package services

import (
	"errors"
	"strings"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/dto"
	"github.com/Wrestler094/shortener/internal/persistence"
	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/utils"
)

type URLService struct {
	storage     storage.IStorage
	fileStorage *persistence.FileStorage
}

func NewURLService(s storage.IStorage, fs *persistence.FileStorage) *URLService {
	return &URLService{storage: s, fileStorage: fs}
}

func (s *URLService) SaveURL(url string) (string, error) {
	originalURL := strings.TrimSpace(url)

	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		return "", errors.New("invalid URL format")
	}

	shortID, err := utils.GenerateShortID()
	if err != nil {
		return "", errors.New("failed to generate short id")
	}

	// TODO: Сделать проверку на случай если id или URL уже существует
	s.storage.Save(shortID, originalURL)
	s.fileStorage.SaveURL(shortID, originalURL)

	return shortID, nil
}

func (s *URLService) GetURLByID(id string) (string, bool) {
	return s.storage.Get(id)
}

func (s *URLService) SaveBatch(batch dto.BatchRequestList) (dto.BatchResponseList, error) {
	var response dto.BatchResponseList

	urls := make(map[string]string) // shortURL[originalURL]

	for _, item := range batch {
		shortID, err := utils.GenerateShortID()
		if err != nil {
			return nil, err
		}

		urls[shortID] = strings.TrimSpace(item.OriginalURL)

		response = append(response, dto.BatchResponseItem{
			CorrelationID: item.CorrelationID,
			ShortURL:      configs.FlagBaseAddr + "/" + shortID,
		})
	}

	// Сохраняем атомарно
	err := s.storage.SaveBatch(urls)
	if err != nil {
		return nil, err
	}

	// Запись в файл
	for shortID, orig := range urls {
		s.fileStorage.SaveURL(shortID, orig)
	}

	return response, nil
}
