package services

import (
	"errors"
	"strings"

	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/storage/file"
	"github.com/Wrestler094/shortener/internal/utils"
)

func SaveURL(url string) (string, error) {
	originalURL := strings.TrimSpace(url)

	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		return "", errors.New("invalid URL format")
	}

	shortID, err := utils.GenerateShortID()
	if err != nil {
		return "", errors.New("failed to generate short id")
	}

	// TODO: Сделать проверку на случай если id или URL уже существует
	storage.Storage.Save(shortID, originalURL)
	file.SaveURL(shortID, originalURL)

	return shortID, nil
}
