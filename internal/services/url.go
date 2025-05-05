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

// URLService предоставляет методы для управления сокращёнными URL.
type URLService struct {
	storage     storage.IStorage         // Основное хранилище URL
	fileStorage *persistence.FileStorage // Файловое хранилище для бэкапа
	deleter     deleter.Deleter          // Сервис для асинхронного удаления URL
}

// NewURLService создаёт и возвращает новый экземпляр URLService.
// s - основное хранилище URL
// fs - файловое хранилище для бэкапа
// dl - сервис для асинхронного удаления URL
func NewURLService(s storage.IStorage, fs *persistence.FileStorage, dl deleter.Deleter) *URLService {
	return &URLService{storage: s, fileStorage: fs, deleter: dl}
}

// SaveURL сохраняет оригинальный URL и генерирует для него короткий идентификатор.
// Если URL уже существует в хранилище, возвращает существующий короткий ID и ошибку.
// url - оригинальный URL для сохранения
// userID - идентификатор пользователя
// Возвращает:
// - сокращенный URL
// - ошибку, если URL уже существует или произошла другая ошибка
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

// SaveBatch сохраняет пакет URL-ов, ассоциированных с пользователем.
// Возвращает список сгенерированных сокращённых URL и correlation ID.
// batch - список URL для сохранения с их correlation ID
// userID - идентификатор пользователя
// Возвращает:
// - список сохраненных URL с их correlation ID
// - ошибку, если произошла ошибка при сохранении
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

// GetURLByID возвращает оригинальный URL по его короткому идентификатору.
// Также указывает, найден ли он и был ли помечен как удалённый.
// id - короткий идентификатор URL
// Возвращает:
// - оригинальный URL
// - флаг удаления
// - флаг наличия URL в хранилище
func (s *URLService) GetURLByID(id string) (string, bool, bool) {
	return s.storage.Get(id)
}

// GetUserURLs возвращает все URL, ранее сохранённые конкретным пользователем.
// uuid - идентификатор пользователя
// Возвращает:
// - список URL пользователя
// - ошибку, если произошла ошибка при получении URL
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

// DeleteUserURLs помещает переданные URL пользователя в очередь на удаление.
// userID - идентификатор пользователя
// urls - список сокращенных URL для удаления
// Возвращает ошибку, если произошла ошибка при добавлении в очередь
func (s *URLService) DeleteUserURLs(userID string, urls []string) error {
	if len(urls) == 0 {
		return nil
	}

	for _, shortID := range urls {
		s.deleter.QueueForDeletion(shortID, userID)
	}

	return nil
}
