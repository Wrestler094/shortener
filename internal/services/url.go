package services

import (
	"errors"
	"strings"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/deleter"
	"github.com/Wrestler094/shortener/internal/dto"
	"github.com/Wrestler094/shortener/internal/persistence"
	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/storage/postgres"
	"github.com/Wrestler094/shortener/internal/utils"
)

type URLService struct {
	storage     storage.IStorage
	fileStorage *persistence.FileStorage
	deleter     deleter.Deleter
}

func NewURLService(s storage.IStorage, fs *persistence.FileStorage, dl deleter.Deleter) *URLService {
	return &URLService{storage: s, fileStorage: fs, deleter: dl}
}

func (s *URLService) SaveURL(url string, userID string) (string, error) {
	original := strings.TrimSpace(url)
	if !strings.HasPrefix(original, "http://") && !strings.HasPrefix(original, "https://") {
		return "", errors.New("invalid URL format")
	}

	// TODO: Сделать проверку на случай если id или URL уже существует
	shortID, err := utils.GenerateShortID()
	if err != nil {
		return "", err
	}

	err = s.storage.Save(shortID, original, userID)
	if err != nil {
		if errors.Is(err, postgres.ErrURLAlreadyExists) {
			existingShort, lookupErr := s.storage.FindShortByOriginalURL(original)
			if lookupErr != nil {
				return "", lookupErr
			}
			return existingShort, postgres.ErrURLAlreadyExists
		}
		return "", err
	}

	s.fileStorage.SaveURL(shortID, original)
	return shortID, nil
}

func (s *URLService) SaveBatch(batch []dto.BatchRequestItem, userID string) ([]dto.BatchResponseItem, error) {
	var response []dto.BatchResponseItem

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
	err := s.storage.SaveBatch(urls, userID)
	if err != nil {
		return nil, err
	}

	// Запись в файл
	for shortID, orig := range urls {
		s.fileStorage.SaveURL(shortID, orig)
	}

	return response, nil
}

func (s *URLService) GetURLByID(id string) (string, bool, bool) {
	return s.storage.Get(id)
}

func (s *URLService) GetUserURLs(uuid string) ([]dto.UserURLItem, error) {
	rawURLs, err := s.storage.GetUserURLs(uuid)
	if err != nil {
		return nil, err
	}

	urls := make([]dto.UserURLItem, 0, len(rawURLs))
	for _, r := range rawURLs {
		urls = append(urls, dto.UserURLItem{
			ShortURL:    configs.FlagBaseAddr + "/" + r.ShortURL,
			OriginalURL: r.OriginalURL,
		})
	}

	return urls, nil
}

func (s *URLService) DeleteUserURLs(userID string, urls []string) error {
	if len(urls) == 0 {
		return nil
	}

	for _, shortID := range urls {
		s.deleter.QueueForDeletion(shortID, userID)
	}

	return nil
}
